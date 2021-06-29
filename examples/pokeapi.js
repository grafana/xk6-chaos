import http from 'k6/http';
import { sleep, check } from 'k6';
import papaparse from 'https://jslib.k6.io/papaparse/5.1.1/index.js';
import { SharedArray } from 'k6/data';
import { podkillers as Podkillers } from 'k6/x/chaos/experiments';
import { Pods } from 'k6/x/chaos/k8s';
import chaos from 'k6/x/chaos';
import { textSummary } from 'https://jslib.k6.io/k6-summary/0.0.1/index.js';

const pod = new Pods();
console.log(`Running simskij/xk6-chaos@${chaos.version}.`);

export const options = {
    scenarios: {
        // pokeapi: {
        //     executor: 'ramping-vus',
        //     exec: 'catchEmAll',
        //     startVUs: 1,
        //     stages: [
        //         { duration: '1m', target: 1 },
        //     ],
        //     gracefulRampDown: '1s',
        // },
        chaos: {
            executor: 'per-vu-iterations',
            exec: 'killPod',
            vus: 1,
            iterations: 1,
            // startTime: '5s',
        },
    },
    thresholds: {
        http_req_failed: ['rate<0.05'],
        http_req_duration: ['p(95)<100'],
    },
};

// export const options = { 
//     scenarios: {
//         pokeapi: {
//             executor: 'per-vu-iterations',
//             exec: 'catchEmAll',
//             vus: 1,
//             iterations: 1,
//         }
//     }
// }

const domain = 'http://cluster.nicolevanderhoeven.com/api/v2';
const sharedData = new SharedArray("Shared Logins", function() {
    let data = papaparse.parse(open('pokemon.csv'), { header: true }).data;
    return data;
});

export function catchEmAll() {
    GetPokemon();
    ThinkTime();
}

export function GetPokemon() {
    // Get random mon from shared array
    let randomMon = sharedData[Math.floor(Math.random() * sharedData.length)];
    // console.log(JSON.stringify(randomMon) + ' selected');

    let res = http.get(domain + '/pokemon/' + randomMon.name, {tags: { name: '01_GetPokemon' }});
    check(res, {
        'is status 200': (r) => r.status === 200,
        '01-text verification': (r) => r.body.includes(randomMon.name)
    });
    sleep(Math.random() * 5);
}

export function ThinkTime() {
    sleep(Math.random() * 5);
}

export function killPod() {

    Podkillers.setStartingPods(pod.list().length);

    let victim = 'nothing';
    console.log('Podkillers ' + JSON.stringify(Podkillers));
    // Iterate through the list of pods to determine which one to kill.
    for (let i = 0; i < pod.list().length; i++) {
        victim = pod.list()[i];
        // console.log('in loop', i, ': victim:', victim);
        // Choose a pod with a name starting with a substring to kill.
        if (victim.startsWith('web')) {
            console.log('Victim chosen:', victim);
            break;
        }
    }
    
    // Kill chosen pod.
    console.log(`Killing pod ${victim}.`);
    // pod.killByName('default', victim);
    Podkillers.killPod('default', victim);
    console.log(`There are currently ${pod.list().length} pods after killing ${victim}.`);

    Podkillers.addVictim(victim);
}

export function generateChaosSummary(data) {
    let victims = Podkillers.getVictims();
    let numPodsBegin = Podkillers.getStartingPods();
    let summary = data;
    console.log('victims: ' + victims);

    let chaosSummary = `
    \n\nxk6-CHAOS
    -----
    EXPERIMENT: Pod termination.
    HYPOTHESIS: After terminating a pod, the application still operates within thresholds.
    
    Number of pods before termination: ${numPodsBegin}
    ${victims.length} pod(s) terminated: ${victims}
    Number of pods at the end of the test: ${pod.list().length}

    `;
    
    return chaosSummary;
}

export function handleSummary(data) {
    
    return {
        'stdout': textSummary(data, { indent: ' ', enableColors: true}) + generateChaosSummary(data),
        // 'raw-data.json': JSON.stringify(data)
    };
}
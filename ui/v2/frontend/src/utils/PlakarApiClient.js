import {createDummySnapshotItems, fetchSnapshotPage} from "./DataGenerator";
import {dummmyFetchConfig, dummyFetchSnapshotPage, dummyFetchSnapshotsPath, snapshots} from "./DemoRepo";

export function fetchConfig(apiUrl) {
    return dummmyFetchConfig();
}

export function fetchSnapshots(apiUrl, page, pageSize) {
    return dummyFetchSnapshotPage(apiUrl, page, pageSize);
}

export function fetchSnapshotsPath(apiUrl, pathId, page, pageSize) {
    return dummyFetchSnapshotsPath(apiUrl, pathId, page, pageSize);
}


function search(searchParams) {
    return [];
}
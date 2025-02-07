# Job Queue Processor - Technical Specification

### Core Requirements

#### Job Types

**Fast Job**

    Processing time: 1 second

    Example task: String transformation

    Priority: High

**Medium Job**

    Processing time: 3 seconds

    Example task: Data aggregation

    Priority: Medium

**Slow Job**

    Processing time: 5 seconds

    Example task: Complex calculation

    Priority: Low

#### Data Structures

Job Interface

```js

interfaceJob{

id:string;           // Unique identifier

type:'fast'|'medium'|'slow';

payload:{

data:string;     // Job input data

};

createdAt:number;    // Unix timestamp

status:'pending'|'processing'|'completed'|'failed';

}

```

Result Interface

```js

interfaceJobResult{

jobId:string;

startTime:number;    // Processing start time

endTime:number;      // Processing end time

duration:number;     // Processing duration in ms

status:'completed'|'failed';

error?:string;      // Error message if failed

}


```

### Functional Requirements

**Queue Management**

-Maintain separate queues for each job type

-Maximum 5 concurrent jobs processing

-Jobs of the same type should not block other types

-FIFO processing within same job type

**Processing Logic**

-Fast jobs: Simple string operations (e.g., reverse, uppercase)

-Medium jobs: Array operations (e.g., sorting, filtering)

-Slow jobs: Complex calculations (e.g., prime numbers)

**Concurrency**

-Use appropriate concurrency primitives:

-Go: goroutines and channels

-Node.js: Worker Threads

-Implement worker pool pattern (if possible)

**Error Handling**

-Failed jobs should not crash the system

-Retry mechanism not required

-Log errors with stack traces

-Mark failed jobs appropriately

**Monitoring**

-Log start and completion of each job

-Track processing time for each job

-Maintain count of completed/failed jobs

-Calculate average processing time per job type

**Graceful Shutdown**

-Handle SIGTERM signal

-Complete processing jobs before shutdown

-Log final statistics before exit

**Example Usage**

```js

constjobs= [

{

        id:"1",

        type:"fast",

        payload:{ data:"reverse this string"}

},

{

        id:"2",

        type:"medium",

        payload:{ data:"[1,5,2,7,3]"}

},

{

        id:"3",

        type:"slow",

        payload:{ data:"50"}// ex : Calculate prime numbers up to 50

}

];

```

```bash

[2024-12-16 10:15:00] Job 1 started - Type: fast

[2024-12-16 10:15:01] Job 1 completed - Duration: 1002ms

[2024-12-16 10:15:00] Job 2 started - Type: medium

[2024-12-16 10:15:03] Job 2 completed - Duration: 3015ms

```

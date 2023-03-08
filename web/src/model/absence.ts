export interface Absence {
    id: number;
    type: string;
    userId: string;
    name: string;
    status: string;
    fromDate: Date;
    toDate: Date;
    creation: Date;
}

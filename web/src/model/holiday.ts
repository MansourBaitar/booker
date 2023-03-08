import { Dayjs } from "dayjs";

export interface Holiday {
    id: number;
    type: string;
    userId: string;
    name: string;
    status: string;
    fromDate: Date | Dayjs | null;
    toDate: Date | Dayjs | null;
    creation: Date;
}

import { Project } from "./project";

export interface ProjectGroup {
    id?: number;
    name: string;
    user_id: number;
    projects?: Project[];
}
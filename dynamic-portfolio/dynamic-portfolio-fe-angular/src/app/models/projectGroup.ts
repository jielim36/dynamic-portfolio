import { Project } from "./project";

export interface ProjectGroup {
    id: number;
    name: string;
    projects: Project[];
}
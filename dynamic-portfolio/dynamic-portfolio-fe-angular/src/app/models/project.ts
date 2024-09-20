import { ProjectImage } from "./projectImage";

export interface Project {
    id: number;
    group_id: number;
    title: string;
    description: string;
    start_date: string;
    end_date?: string;
    project_images: ProjectImage[];
}
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ProjectGroup } from '../models/projectGroup';
import { HttpResponse } from '../models/response';

@Injectable({
  providedIn: 'root',
})
export class ProjectService {

  constructor(private readonly http: HttpClient) { }


  getProjectGroupsByUserId(userId: number) {
    return this.http.get<HttpResponse<ProjectGroup[]>>(`/groups/user/${userId}`)
  }

  addProjectGroup(projectGroup: ProjectGroup) {
    return this.http.post<HttpResponse<number>>("/groups", projectGroup)
  }
}

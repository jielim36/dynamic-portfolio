import { Component } from '@angular/core';
import { ProjectGroup } from '../../models/projectGroup';
import { ProjectService } from '../../services/project.service';
import { HttpResponse } from '../../models/response';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [
    CommonModule
  ],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {

  projectGroups: ProjectGroup[] = []
  currentGroup: ProjectGroup | undefined;

  constructor(private readonly projectService: ProjectService) {
    this.init();
  }

  init() {
    this.projectService.getProjectGroupsByUserId(1).subscribe(
      (response: HttpResponse<ProjectGroup[]>) => {
        if (response.data) {
          this.projectGroups = response.data;
          this.currentGroup = response.data[0];
          console.log(this.projectGroups);

        }
      },
      (error) => {
        console.error('Error fetching projects:', error);
      }
    );
  }

  changeGroup(groupId: number) {
    this.currentGroup = this.projectGroups.find((grp) => grp.id === groupId)

  }

}

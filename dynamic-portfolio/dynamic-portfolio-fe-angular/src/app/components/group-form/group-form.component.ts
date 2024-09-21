import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { ProjectService } from '../../services/project.service';
import { ProjectGroup } from '../../models/projectGroup';
import { HttpResponse } from '../../models/response';

@Component({
  selector: 'app-group-form',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule
  ],
  templateUrl: './group-form.component.html',
  styleUrl: './group-form.component.css'
})
export class GroupFormComponent {
  projectGroupForm: FormGroup;

  constructor(private fb: FormBuilder, private readonly projectService: ProjectService) {
    this.projectGroupForm = this.fb.group({
      name: ['', [Validators.required, Validators.minLength(3)]],
      user_id: [1, [Validators.required]]
    });
  }

  onSubmit() {
    if (this.projectGroupForm.valid) {
      console.log('Form Submitted', this.projectGroupForm.value);
      const projectGroup: ProjectGroup = {
        name: this.projectGroupForm.value.name,
        user_id: this.projectGroupForm.value.user_id
      };
      const res = this.projectService.addProjectGroup(projectGroup)
      res.subscribe(
        (response: HttpResponse<number>) => {
          console.log(response);
        },
        (error) => {
          console.error('Error Add Project Group:', error);
        }
      )

    } else {
      console.log('Form is invalid');
    }
  }
}

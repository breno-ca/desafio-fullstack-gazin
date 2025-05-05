import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { NiveisRoutingModule } from './niveis-routing.module';
import { NiveisListComponent } from './niveis-list/niveis-list.component';
import { DataTableComponent } from 'src/app/shared/components/data-table/data-table.component';
import { NiveisFormComponent } from './niveis-form/niveis-form.component';
import { ReactiveFormsModule } from '@angular/forms';
import { FormComponent } from 'src/app/shared/components/form/form.component';

@NgModule({
  declarations: [NiveisListComponent, NiveisFormComponent],
  imports: [CommonModule, NiveisRoutingModule, ReactiveFormsModule, DataTableComponent, FormComponent],
})
export class NiveisModule {}

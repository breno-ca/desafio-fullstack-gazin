import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { DesenvolvedoresRoutingModule } from './desenvolvedores-routing.module';
import { DesenvolvedoresFormComponent } from './desenvolvedores-form/desenvolvedores-form.component';
import { DesenvolvedoresListComponent } from './desenvolvedores-list/desenvolvedores-list.component';
import { ReactiveFormsModule } from '@angular/forms';
import { DataTableComponent } from 'src/app/shared/components/data-table/data-table.component';
import { FormComponent } from 'src/app/shared/components/form/form.component';

@NgModule({
  declarations: [DesenvolvedoresListComponent, DesenvolvedoresFormComponent],
  imports: [CommonModule, DesenvolvedoresRoutingModule, ReactiveFormsModule, DataTableComponent, FormComponent],
})
export class DesenvolvedoresModule {}

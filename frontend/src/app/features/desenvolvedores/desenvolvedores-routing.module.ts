import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DesenvolvedoresListComponent } from './desenvolvedores-list/desenvolvedores-list.component';

const routes: Routes = [
  {
    path: '',
    component: DesenvolvedoresListComponent,
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class DesenvolvedoresRoutingModule { }

import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { NiveisListComponent } from './niveis-list/niveis-list.component';

const routes: Routes = [
  {
    path: '',
    component: NiveisListComponent,
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class NiveisRoutingModule {}

import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DashboardComponent } from './dashboard.component';

const routes: Routes = [
  {
    path: '',
    component: DashboardComponent,
    children: [
      {
        path: 'niveis',
        loadChildren: () => import('../features/niveis/niveis.module').then((m) => m.NiveisModule),
      },
      {
        path: 'desenvolvedores',
        loadChildren: () => import('../features/desenvolvedores/desenvolvedores.module').then((m) => m.DesenvolvedoresModule),
      },
    ],
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class DashboardRoutingModule {}

import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { DataTableComponent } from './shared/components/data-table/data-table.component';
import { ToastComponent } from './shared/components/toast/toast.component';
import { ModalComponent } from './shared/components/modal/modal.component';

@NgModule({
  declarations: [AppComponent],
  imports: [BrowserModule, AppRoutingModule, HttpClientModule, DataTableComponent, ToastComponent, ModalComponent],
  bootstrap: [AppComponent],
})
export class AppModule {}

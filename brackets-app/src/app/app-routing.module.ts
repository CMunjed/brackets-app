import { NgModule } from '@angular/core';
import { RouterModule, ROUTES, Routes } from '@angular/router';
import {MatSlideToggleModule} from '@angular/material/slide-toggle';

const routes: Routes = [];

@NgModule({
  imports: [
    RouterModule,
    MatSlideToggleModule
  ],
  exports: [
    RouterModule,
    MatSlideToggleModule
  ]
})
export class AppRoutingModule { }

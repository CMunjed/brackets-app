import { NgModule } from '@angular/core';
import { RouterModule, ROUTES, Routes } from '@angular/router';
import { AppComponent } from './app.component';
import { SignInComponent } from './sign-in/sign-in.component';
import {MatSlideToggleModule} from '@angular/material/slide-toggle';

const routes: Routes = [


];

@NgModule({
  imports: [
    RouterModule.forRoot(routes),
    RouterModule,
    MatSlideToggleModule
  ],
  exports: [
    RouterModule,
    MatSlideToggleModule
  ]
})
export class AppRoutingModule { }

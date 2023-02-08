import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import {MatSlideToggleModule} from '@angular/material/slide-toggle';
import {MatTableModule} from '@angular/material/table';


import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

@NgModule({
  declarations: [
    AppComponent
  ],

  imports: [
    BrowserModule,
    AppRoutingModule,
    MatSlideToggleModule,
    MatTableModule
     
    
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

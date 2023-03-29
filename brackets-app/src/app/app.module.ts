import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import {MatSlideToggleModule} from '@angular/material/slide-toggle';
import {MatTableModule} from '@angular/material/table';


import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { HttpClientModule } from '@angular/common/http';

import { CreateBracketService } from './create-bracket.service';


@NgModule({
  declarations: [
    AppComponent,
   
  ],

  imports: [
    BrowserModule,
    AppRoutingModule,
    MatSlideToggleModule,
    MatTableModule,
    HttpClientModule,
    BrowserAnimationsModule
     
    
  ],
  providers: [
   
    CreateBracketService
  ],
  bootstrap: [AppComponent]
  
})
export class AppModule { }

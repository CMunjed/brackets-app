import { Component } from '@angular/core';

export interface PeriodicElement {
  name: string;
  position: number;
  score: number;
  
}
const ELEMENT_DATA: PeriodicElement[] = [
  {position: 1, name: 'Gators', score: 10079},
  {position: 2, name: 'Bucks', score: 400.26},
  {position: 3, name: 'Packers', score: 69.41},
  {position: 4, name: 'Blue Devils', score: 9.0122},
  {position: 5, name: 'Honey Badgers', score: 1.0811},
  
];

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'brackets-app';
  displayedColumns: string[] = ['position', 'name', 'weight'];
  dataSource = ELEMENT_DATA;
  
}

import { Component, OnInit, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { delay } from 'rxjs/operators';
import { EventEmitter } from '@angular/core';
import { HttpClient } from '@angular/common/http';



@Component({
  selector: 'app-signin-dialog',
  templateUrl: './signin-dialog.component.html',
  styleUrls: ['./signin-dialog.component.css']
})
export class SigninDialogComponent {

  isLoggedIn: boolean = false;
  onSignInSuccess = new EventEmitter();
  email: string = 'conner123@gmail.com';
  password: string = 'Ihopethisworks';

  constructor(
    public dialogRef: MatDialogRef<SigninDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: any,
    private snackBar: MatSnackBar,
    private http: HttpClient

  ) { 
  
  }

  ngOnInit(): void {  
  }

  onSubmit(email: string, password: string): void {
    
    this.isLoggedIn = true;
    this.dialogRef.close();
    //this.snackBar.open('Signed in successfully!', 'Close', { duration: 30000 });
  
    this.http.post('http://localhost:3000/users/signup', { email, password }).subscribe(
      (response) => {
        // Handle successful login
        console.log('Login successful:', response);
      },
      (error) => {
        // Handle login error
        console.error('Login error:', error);
      }
    );

        console.log(email);
  }


  onCancel(): void {
    this.dialogRef.close();
  }
}



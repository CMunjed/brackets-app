import { Component, OnInit, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { delay } from 'rxjs/operators';
import { EventEmitter } from '@angular/core';


@Component({
  selector: 'app-signin-dialog',
  templateUrl: './signin-dialog.component.html',
  styleUrls: ['./signin-dialog.component.css']
})
export class SigninDialogComponent {

  isLoggedIn: boolean = false;
  onSignInSuccess = new EventEmitter();

  constructor(
    public dialogRef: MatDialogRef<SigninDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: any,
    private snackBar: MatSnackBar
  ) { 
  
  }

  ngOnInit(): void {  
  }

  onSubmit(): void {
    
    this.isLoggedIn = true;
    this.dialogRef.close();
    this.snackBar.open('Signed in successfully!', 'Close', { duration: 30000 });
  
  }


  onCancel(): void {
    this.dialogRef.close();
  }
}



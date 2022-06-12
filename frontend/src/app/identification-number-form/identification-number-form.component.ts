import { HttpErrorResponse } from '@angular/common/http';
import { AfterViewInit, Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { catchError, finalize, of } from 'rxjs';
import { IdentificationNumber } from '../shared/identification-number.model';
import { IdentificationNumberService } from '../shared/identification-number.service';

@Component({
  selector: 'app-identification-number-form',
  templateUrl: './identification-number-form.component.html',
  styleUrls: ['./identification-number-form.component.css']
})
export class IdentificationNumberFormComponent implements AfterViewInit {

  public form: FormGroup;

  constructor(
      private identificationNumberService: IdentificationNumberService,
      private snackBar: MatSnackBar) { 
    this.form = new FormGroup({
      number: new FormControl('')
    })
  }

  ngAfterViewInit(): void {
    this.form.reset();
  }

  get serverError():string|null {
    let control = this.form.get("number");
    if(control &&  control.errors){
      return control.errors['serverError'];
    }
    return null;
  }

  save():void{
    this.identificationNumberService.create(this.form.value)
      .pipe(
        catchError((res) => {
          let control = this.form.get("number");
          if(control){
            control.setErrors({
              serverError:res.error[0]
            })
          }     
          return of(null);
        }),
        finalize(() => console.info("Finalize"))
      )
      .subscribe((data:IdentificationNumber|null)=>{
        if(data){
          this.snackBar.open("Identification number saved!","Close",{
            duration: 5000
          });
          this.form.reset();
        }
      });
  }
}

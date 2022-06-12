import { AfterViewInit, Component } from '@angular/core';
import { MatCheckboxChange } from '@angular/material/checkbox';
import { IdentificationNumberDataSource } from '../shared/identification-number.datasource';
import { IdentificationNumber } from '../shared/identification-number.model';
import { IdentificationNumberService } from '../shared/identification-number.service';

@Component({
  selector: 'app-identification-number-list',
  templateUrl: './identification-number-list.component.html',
  styleUrls: ['./identification-number-list.component.css']
})
export class IdentificationNumberListComponent implements AfterViewInit {

  displayedColumns: string[] = ['id','number', 'blocked'];

  constructor(private identificationNumberService:IdentificationNumberService,private identificationNumberDataSource: IdentificationNumberDataSource) {}
  
  ngAfterViewInit(): void {
    this.identificationNumberDataSource.init();
  }

  get dataSource(){
    return this.identificationNumberDataSource;
  }

  get form(){
    return this.identificationNumberDataSource.form;
  }

  toggle(event:MatCheckboxChange,identificationNumber:IdentificationNumber):void{
    identificationNumber.blocked = event.checked;
    this.identificationNumberService.update(identificationNumber)
        .subscribe((data:IdentificationNumber)=>{
          console.info("IdentificationNumber updated!");
        });
  }
}

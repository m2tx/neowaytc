import { AfterViewInit, Component, OnInit } from '@angular/core';
import { MatCheckboxChange } from '@angular/material/checkbox';
import { IdentificationNumberDataSource } from '../shared/identification-number.datasource';
import { IdentificationNumber } from '../shared/identification-number.model';
import { IdentificationNumberService } from '../shared/identification-number.service';

@Component({
  selector: 'app-identification-number-list',
  templateUrl: './identification-number-list.component.html',
  styleUrls: ['./identification-number-list.component.css']
})
export class IdentificationNumberListComponent implements OnInit {

  displayedColumns: string[] = ['id','number', 'blocked'];
  public dataSource:IdentificationNumberDataSource;

  constructor(
        private identificationNumberService:IdentificationNumberService) {  
        this.dataSource = new IdentificationNumberDataSource(identificationNumberService); 
  }

  ngOnInit(): void { 
    this.dataSource.init();
  }

  toggle(event:MatCheckboxChange,identificationNumber:IdentificationNumber):void{
    identificationNumber.blocked = event.checked;
    this.identificationNumberService.update(identificationNumber)
        .subscribe((data:IdentificationNumber)=>{
          console.info("IdentificationNumber updated!");
        });
  }
}

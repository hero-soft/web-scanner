import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CallComponent } from './call/call.component';
import { MatButtonModule } from '@angular/material/button';
import {MatTooltipModule} from '@angular/material/tooltip';


@NgModule({
  declarations: [
    CallComponent
  ],
  imports: [
    CommonModule,
    MatButtonModule,
    MatTooltipModule,
  ],
  exports: [
    CallComponent
  ]
})
export class CallModule { }

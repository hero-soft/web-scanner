import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PlayerService } from './player.service';
import { AutoplayDialogComponent } from './autoplay-dialog/autoplay-dialog.component';
import {MatDialogModule} from '@angular/material/dialog';


@NgModule({
  declarations: [
    AutoplayDialogComponent
  ],
  imports: [
    CommonModule,
    MatDialogModule,
  ],
})
export class PlayerModule { }

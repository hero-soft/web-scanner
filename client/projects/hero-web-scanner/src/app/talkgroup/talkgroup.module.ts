import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { TalkgroupDialogComponent } from './dialog/dialog.component';
import { MatDialogModule } from '@angular/material/dialog';
import { MatButtonModule } from '@angular/material/button';
import { HttpClientModule } from '@angular/common/http';
import { TalkgroupListComponent } from './list/list.component';


@NgModule({
  declarations: [
    TalkgroupDialogComponent,
    TalkgroupListComponent
  ],
  imports: [
    CommonModule,
    HttpClientModule,
    MatDialogModule,
    MatButtonModule,
  ],
  exports: [
    TalkgroupDialogComponent
  ]
})
export class TalkgroupModule { }

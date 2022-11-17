import { NgModule, isDevMode } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { WebsocketService } from './websocket.service';
import { TgLookupPipe } from './tg-lookup.pipe';
import { CallModule } from './call/call.module';
import { PlayerModule } from './player/player.module';
import { SettingsModule } from './settings/settings.module';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatTooltipModule} from '@angular/material/tooltip';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { TalkgroupModule } from './talkgroup/talkgroup.module';

import { DBConfig, NgxIndexedDBModule } from 'ngx-indexed-db';
import { AngularRuntimeConfigModule } from 'angular-runtime-config';
import { Settings } from './settings/settings.type';
import { ServiceWorkerModule } from '@angular/service-worker';


const dbConfig: DBConfig  = {
  name: 'HeroWebScanner',
  version: 1,
  objectStoresMeta: [{
    store: 'settings',
    storeConfig: { keyPath: 'id', autoIncrement: true },
    storeSchema: [
      { name: 'settings', keypath: 'settings', options: { unique: false } },
    ]
  }]
};

@NgModule({
  declarations: [
    AppComponent,
    TgLookupPipe
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    NgxIndexedDBModule.forRoot(dbConfig),
    AngularRuntimeConfigModule.forRoot(Settings),

    CallModule,
    PlayerModule,
    SettingsModule,
    BrowserAnimationsModule,
    MatToolbarModule,
    MatIconModule,
    MatButtonModule,
    MatSidenavModule,
    MatTooltipModule,
    MatProgressSpinnerModule,
    TalkgroupModule,
    ServiceWorkerModule.register('ngsw-worker.js', {
      enabled: !isDevMode(),
      // Register the ServiceWorker as soon as the application is stable
      // or after 30 seconds (whichever comes first).
      registrationStrategy: 'registerWhenStable:30000'
    }),
  ],
  providers: [
    WebsocketService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }

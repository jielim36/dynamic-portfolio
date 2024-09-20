import { mergeApplicationConfig, ApplicationConfig } from '@angular/core';
import { provideServerRendering } from '@angular/platform-server';
import { provideHttpClient, withInterceptors } from "@angular/common/http";
import { appConfig } from './app.config';
import { appHttpInterceptor } from './app.http.interceptor';

const serverConfig: ApplicationConfig = {
  providers: [
    provideServerRendering(),
    provideHttpClient(withInterceptors([appHttpInterceptor]))
  ]
};

export const config = mergeApplicationConfig(appConfig, serverConfig);

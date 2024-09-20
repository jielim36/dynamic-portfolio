import { HttpInterceptorFn } from "@angular/common/http";
import { inject } from "@angular/core";
import { env } from "process";
import { environment } from "./environments/environment";
// import AppAuthService from "./services/auth.service";

export const appHttpInterceptor: HttpInterceptorFn = (request, next) => {
    // const appAuthService = inject(AppAuthService);
    const SERVER_API_BASE_URL = environment.apiUrl;

    request = request.clone({
        // setHeaders: { Authorization: `Bearer ${appAuthService.token()}` }
        url: `${SERVER_API_BASE_URL}${request.url}`
    });

    return next(request);
};
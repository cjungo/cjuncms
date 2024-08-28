import { ElMessage } from "element-plus";
import { debounceTime, Subject } from "rxjs";
import { type AppContext } from "vue";

export const tipOk = new Subject<string>();
export const tipInfo = new Subject<string>();
export const tipWarn = new Subject<string>();
export const tipError = new Subject<Error>();

const subscribe = (
  subject: Subject<string>,
  type: "success" | "warning" | "info",
  appContext: AppContext
) => {
  subject.subscribe({
    next: (message) => {
      ElMessage(
        {
          message: <p>{message}</p>,
          grouping: true,
          type: type,
        },
        appContext
      );
    },
  });
};

export const tipSubscribe = (appContext: AppContext) => {
  subscribe(tipOk, "success", appContext);
  subscribe(tipInfo, "info", appContext);
  subscribe(tipWarn, "warning", appContext);

  tipError.pipe(debounceTime(500)).subscribe({
    next: (error: any) => {
      let message = error.response?.data?.message;
      if (!message) {
        message = error.message;
      }
      ElMessage(
        {
          message: <p>{message}</p>,
          grouping: true,
          type: "error",
        },
        appContext
      );
    },
  });
};

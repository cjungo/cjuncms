import { ElMessage } from "element-plus";
import { Subject } from "rxjs";
import { type AppContext } from "vue";

export const tipOk = new Subject<string>();
export const tipInfo = new Subject<string>();
export const tipWarn = new Subject<string>();
export const tipError = new Subject<Error>();

export const tipSubscribe = (appContext: AppContext) => {
  tipOk.subscribe({
    next: (message) => {
      ElMessage(
        {
          message: <p>{message}</p>,
          grouping: true,
          type: "success",
        },
        appContext
      );
    },
  });

  tipError.subscribe({
    next: (error) => {
      console.log("next", error, appContext);
      ElMessage(
        {
          message: <p>{error.message}</p>,
          grouping: true,
          type: "error",
        },
        appContext
      );
    },
  });
};

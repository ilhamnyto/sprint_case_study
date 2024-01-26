import { CheckIcon, TrashIcon } from "@radix-ui/react-icons";
import TodoDialog from "./TodoDialog";
import { completeTask, deleteTask } from "@/lib/utils";
import { AppContext } from "@/context/appContext";
import { useContext } from "react";

export default function TodoAction({ data, isSubTask }) {
  const { state, dispatch } = useContext(AppContext);

  if (data.completed_at) {
    return (
      <div
        className="bg-red-300 p-1 rounded hover:bg-red-400 transition-all"
        onClick={async () => {
          await deleteTask(data.id, isSubTask);
          isSubTask
            ? dispatch({ type: "DELETE_SUBTASK", payload: data })
            : dispatch({ type: "DELETE_TASK", payload: data });
        }}
      >
        <TrashIcon className="text-white" />
      </div>
    );
  }

  return (
    <div className="flex gap-2 items-center">
      <div
        className="bg-green-300 p-1 rounded hover:bg-green-500 transition-all"
        onClick={async () => {
          await completeTask(data.id, isSubTask);
          isSubTask
            ? dispatch({ type: "COMPLETE_SUBTASK", payload: data })
            : dispatch({ type: "COMPLETE_TASK", payload: data });
        }}
      >
        <CheckIcon className="text-white" />
      </div>
      {!isSubTask && <TodoDialog type="ADD_SUBTASK" data={data} />}
      <TodoDialog
        type={isSubTask ? "UPDATE_SUBTASK" : "UPDATE_TASK"}
        data={data}
      />
      <div
        className="bg-red-300 p-1 rounded hover:bg-red-400 transition-all"
        onClick={async () => {
          await deleteTask(data.id, isSubTask);
          isSubTask
            ? dispatch({ type: "DELETE_SUBTASK", payload: data })
            : dispatch({ type: "DELETE_TASK", payload: data });
        }}
      >
        <TrashIcon className="text-white" />
      </div>
    </div>
  );
}

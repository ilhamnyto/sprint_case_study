"use client";

import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Pencil1Icon, PlusIcon } from "@radix-ui/react-icons";
import { useContext, useState } from "react";
import CalendarInput from "../TaskHeader/CalendarInput";
import { addSubTask, updateTaskOrSubtTask } from "@/lib/utils";
import { formatISO } from "date-fns";
import { DialogClose } from "@radix-ui/react-dialog";
import { AppContext } from "@/context/appContext";

export default function TodoDialog({ type, data }) {
  const { state, dispatch } = useContext(AppContext);

  let initialTodo =
    type == "UPDATE_TASK" || type == "UPDATE_SUBTASK" ? data?.title : "";
  let initialDate =
    type == "UPDATE_TASK" || type == "UPDATE_SUBTASK" ? data?.deadline : "";

  const [date, setDate] = useState(initialDate);
  const [todo, setTodo] = useState(initialTodo);

  const submitHandler = async (event) => {
    event.preventDefault();
    let deadline = date ? formatISO(date, "yyyy-MM-dd'T'HH:mm:ss'Z'") : null;
    switch (type) {
      case "ADD_SUBTASK":
        const res = await addSubTask(data?.id, todo, deadline);
        dispatch({ type: "ADD_SUBTASK", payload: res.data });
        break;
      case "UPDATE_TASK":
        await updateTaskOrSubtTask(data?.id, todo, deadline, false);
        dispatch({
          type: "UPDATE_TASK",
          payload: { id: data?.id, title: todo, deadline: deadline },
        });
        break;
      case "UPDATE_SUBTASK":
        await updateTaskOrSubtTask(data?.id, todo, deadline, true);
        dispatch({
          type: "UPDATE_SUBTASK",
          payload: {
            id: data?.id,
            title: todo,
            deadline: deadline,
            task_id: data.task_id,
          },
        });
        break;
    }
  };
  const changeHandler = (event) => {
    setTodo(event.target.value);
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <div className="bg-slate-300 p-1 rounded hover:bg-gray-700 transition-all">
          {type == "ADD_SUBTASK" && <PlusIcon className="text-white" />}
          {(type == "UPDATE_SUBTASK" || type == "UPDATE_TASK") && (
            <Pencil1Icon className="text-white" />
          )}
        </div>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>
            {type == "ADD_SUBTASK" && "Add Subtask"}
            {type == "UPDATE_SUBTASK" && "Update Subtask"}
            {type == "UPDATE_TASK" && "Update Task"}
          </DialogTitle>
        </DialogHeader>
        <form onSubmit={submitHandler}>
          <div className="space-y-2">
            <div className="">
              <input
                type="text"
                className="w-full border border-gray-200 px-3 py-2 rounded-md text-sm"
                placeholder="Task todo..."
                onChange={changeHandler}
                value={todo}
              />
            </div>
            <div>
              <CalendarInput date={date} setDate={setDate} />
            </div>
          </div>
          <DialogFooter>
            <DialogClose asChild>
              <Button type="submit" className="mt-3">
                {type == "ADD_SUBTASK" && "Add"}
                {(type == "UPDATE_SUBTASK" || type == "UPDATE_TASK") &&
                  "Update"}
              </Button>
            </DialogClose>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}

"use client";

import { useContext, useState } from "react";
import { Button } from "@/components/ui/button";
import CalendarInput from "@/components/TaskHeader/CalendarInput";
import { formatISO } from "date-fns";
import { addTask } from "@/lib/utils";
import { AppContext } from "@/context/appContext";

export default function TaskInput() {
  const [date, setDate] = useState();
  const [todo, setTodo] = useState("");
  const { state, dispatch } = useContext(AppContext);

  const submitHandler = async (event) => {
    event.preventDefault();

    const res = await addTask({
      title: todo,
      deadline: date ? formatISO(date, "yyyy-MM-dd'T'HH:mm:ss'Z'") : null,
    });
    dispatch({ type: "ADD_TASK", payload: res.data });
    setTodo("");
    setDate(null);
  };
  const changeHandler = (event) => {
    setTodo(event.target.value);
  };
  return (
    <div>
      <form className="space-y-2" onSubmit={submitHandler}>
        <input
          type="text"
          className="w-full border border-gray-200 px-3 py-2 rounded-md text-sm"
          placeholder="Task todo..."
          onChange={changeHandler}
          value={todo}
        />
        <CalendarInput date={date} setDate={setDate} />
        <div className="flex justify-end">
          <Button>Add Task</Button>
        </div>
      </form>
    </div>
  );
}

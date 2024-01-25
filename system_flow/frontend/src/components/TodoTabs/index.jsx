"use client";

import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import {
  CheckIcon,
  TrashIcon,
  PlusIcon,
  Pencil1Icon,
} from "@radix-ui/react-icons";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";

import { useState } from "react";
import { CalendarInput } from "../TaskInput";
const { formatISO } = require("date-fns");

export default function TodoTabs() {
  const [todo, setTodo] = useState("");
  const [date, setDate] = useState();
  const [isEditData, setIsEditData] = useState(false);
  const [isAddData, setIsAddData] = useState(false);
  const [datas, setDatas] = useState([
    {
      id: 1,
      title: "title 1",
      created_at: new Date(),
      deadline: null,
      completed_at: null,
      subtask: [
        {
          id: 1,
          task_id: 1,
          title: "subtitle 1",
          created_at: new Date(),
          deadline: null,
          completed_at: null,
        },
        {
          id: 2,
          task_id: 1,
          title: "subtitle 2",
          created_at: new Date(),
          deadline: null,
          completed_at: new Date(),
        },
        {
          id: 3,
          task_id: 1,
          title: "subtitle 3",
          created_at: new Date(),
          deadline: null,
          completed_at: new Date(),
        },
      ],
    },
    {
      id: 2,
      title: "title 2",
      created_at: new Date(),
      deadline: "2024-01-26T00:00:00",
      completed_at: null,
      subtask: [],
    },
    {
      id: 3,
      title: "title 3",
      created_at: new Date(),
      deadline: null,
      completed_at: new Date(),
      subtask: [],
    },
  ]);

  const submitHandler = (event) => {
    event.preventDefault();
    if (date) {
      console.log(formatISO(date));
    }
    console.log(todo, date);
  };
  const changeHandler = (event) => {
    setTodo(event.target.value);
  };

  return (
    <Tabs defaultValue="ongoing" className="w-full">
      <TabsList className="grid w-full grid-cols-2 mb-6">
        <TabsTrigger value="ongoing" className="py-3">
          Ongoing
        </TabsTrigger>
        <TabsTrigger value="completed" className="py-3">
          Completed
        </TabsTrigger>
      </TabsList>
      <div>
        <TabsContent value="ongoing">
          <div className="space-y-2">
            {datas
              .filter((data) => data.completed_at == null)
              .map((data) => (
                <div className="bg-white shadow rounded p-3 flex flex-col cursor-pointer ">
                  {data.subtask.length > 0 ? (
                    <Accordion type="single" collapsible className="w-full">
                      <AccordionItem value="item-1" className="w-full">
                        <div className="flex items-center w-full justify-between">
                          <div className="flex-1">
                            <AccordionTrigger className="md:pr-52">
                              <div className="flex flex-col items-start">
                                {isEditData ? (
                                  <form className="flex justify-start">
                                    <input
                                      type="text"
                                      name=""
                                      id=""
                                      value={data.title}
                                      className="border border-gray-200 px-2 py-1 w-3/4"
                                    />
                                  </form>
                                ) : (
                                  <span>{data.title}</span>
                                )}
                                {data.subtask.length > 0 && (
                                  <div className="flex gap-2">
                                    <span className="text-[10px] mt-1 bg-slate-200 rounded-md px-2 text-slate-800">
                                      {data.subtask.length} subtask
                                    </span>
                                    <span className="text-[10px] mt-1 bg-slate-200 rounded-md px-2 text-slate-800">
                                      {(
                                        (data.subtask.filter(
                                          (t) => t.completed_at != null
                                        ).length /
                                          data.subtask.length) *
                                        100
                                      ).toFixed(0)}
                                      % completed
                                    </span>
                                  </div>
                                )}
                              </div>
                            </AccordionTrigger>
                          </div>
                          <div className="flex gap-2 items-center">
                            <div className="bg-green-300 p-1 rounded hover:bg-green-500 transition-all">
                              <CheckIcon className="text-white" />
                            </div>
                            <div
                              className="bg-slate-300 p-1 rounded hover:bg-gray-700 transition-all"
                              onClick={() => setIsAddData((prev) => !prev)}
                            >
                              <PlusIcon className="text-white" />
                            </div>
                            <div
                              className="bg-slate-300 p-1 rounded hover:bg-gray-700 transition-all"
                              onClick={() => setIsEditData((prev) => !prev)}
                            >
                              <Pencil1Icon className="text-white" />
                            </div>
                            <div className="bg-red-300 p-1 rounded hover:bg-red-400 transition-all">
                              <TrashIcon className="text-white" />
                            </div>
                          </div>
                        </div>

                        <AccordionContent>
                          <div className="space-y-2">
                            {data.subtask.map((el) => (
                              <div className="flex items-center justify-between w-full p-1 rounded line-through">
                                <div className="flex flex-col items-start">
                                  {isEditData ? (
                                    <form>
                                      <input
                                        type="text"
                                        name=""
                                        id=""
                                        value={el.title}
                                      />
                                    </form>
                                  ) : (
                                    <span>{el.title}</span>
                                  )}
                                </div>
                                <div className="flex gap-2 items-center">
                                  <div className="bg-green-300 p-1 rounded hover:bg-green-500 transition-all">
                                    <CheckIcon className="text-white" />
                                  </div>
                                  <div
                                    className="bg-slate-300 p-1 rounded hover:bg-gray-700 transition-all"
                                    onClick={() =>
                                      setIsEditData((prev) => !prev)
                                    }
                                  >
                                    <Pencil1Icon className="text-white" />
                                  </div>
                                  <div className="bg-red-300 p-1 rounded hover:bg-red-400 transition-all">
                                    <TrashIcon className="text-white" />
                                  </div>
                                </div>
                              </div>
                            ))}
                          </div>
                        </AccordionContent>
                      </AccordionItem>
                    </Accordion>
                  ) : (
                    <div className="flex items-center justify-between">
                      <div className="flex flex-col justify-start">
                        <span>{data.title}</span>
                        <div>
                          {data.deadline &&
                            new Date(data.deadline) < new Date() && (
                              <span className="text-[10px] mt-1 px-2 bg-red-300 rounded-md p-1">
                                Overdue
                              </span>
                            )}
                        </div>
                      </div>
                      <div className="flex gap-2 items-center">
                        <div className="bg-green-300 p-1 rounded hover:bg-green-500 transition-all">
                          <CheckIcon className="text-white" />
                        </div>
                        <div
                          className="bg-slate-300 p-1 rounded hover:bg-gray-700 transition-all"
                          onClick={() => setIsAddData((prev) => !prev)}
                        >
                          <PlusIcon className="text-white" />
                        </div>
                        <div
                          className="bg-slate-300 p-1 rounded hover:bg-gray-700 transition-all"
                          onClick={() => setIsEditData((prev) => !prev)}
                        >
                          <Pencil1Icon className="text-white" />
                        </div>
                        <div className="bg-red-300 p-1 rounded hover:bg-red-400 transition-all">
                          <TrashIcon className="text-white" />
                        </div>
                      </div>
                    </div>
                  )}

                  {isAddData && (
                    <div className="mt-3">
                      <form className="space-y-2" onSubmit={submitHandler}>
                        <div className="flex gap-2">
                          <input
                            type="text"
                            className="w-full border border-gray-200 px-3 py-1 rounded-md text-sm"
                            placeholder="Subtask todo..."
                            onChange={changeHandler}
                            value={todo}
                          />
                          <CalendarInput date={date} setDate={setDate} />
                        </div>
                        <div className="flex justify-end">
                          <button className="text-sm bg-black hover:black/80 px-3 py-1 rounded-sm text-white">
                            Add
                          </button>
                        </div>
                      </form>
                    </div>
                  )}
                </div>
              ))}
          </div>
        </TabsContent>
        <TabsContent value="completed">
          <div className="space-y-2">
            {datas
              .filter((data) => data.completed_at != null)
              .map((data) => (
                <div className="bg-white shadow rounded p-3">{data.title}</div>
              ))}
          </div>
        </TabsContent>
      </div>
    </Tabs>
  );
}

function Todo() {}

function TodoList() {}

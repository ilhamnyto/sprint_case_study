"use client";

import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import TodoAction from "./TodoAction";

export default function Todo({ data, isEditData }) {
  return (
    <div
      key={data.id}
      className="bg-white shadow rounded p-3 flex flex-col cursor-pointer "
    >
      {data.subtasks.length > 0 ? (
        <Accordion type="single" collapsible className="w-full">
          <AccordionItem value="item-1" className="w-full">
            <div className="flex items-center w-full justify-between">
              <div className="flex-1">
                <AccordionTrigger className="md:w-full">
                  <div className="flex flex-col items-start">
                    <div className="flex flex-col">
                      <span className="font-bold text-lg">{data.title}</span>
                      <span className="opacity-60 text-[9px]">
                        {data.deadline
                          ? `deadline : ${new Date(
                              data.deadline
                            ).toDateString()}`
                          : "No deadline"}
                      </span>
                    </div>
                    {data.subtasks.length > 0 && (
                      <div className="flex gap-2 w-full items-center">
                        <span className="text-[10px] mt-1 bg-slate-200 rounded-md px-2 py-1 text-slate-800">
                          {data.subtasks.length} subtasks
                        </span>
                        <span className="text-[10px] mt-1 bg-slate-200 rounded-md py-1 px-2 text-slate-800">
                          {(
                            (data.subtasks.filter((t) => t.completed_at != null)
                              .length /
                              data.subtasks.length) *
                            100
                          ).toFixed(0)}
                          % completed
                        </span>
                        {data.deadline &&
                          new Date(data.deadline) < new Date() && (
                            <span className="text-[10px] mt-1 px-2 bg-red-300 rounded-md p-1">
                              Overdue
                            </span>
                          )}
                      </div>
                    )}
                  </div>
                </AccordionTrigger>
              </div>
              <TodoAction data={data} isSubTask={false} />
            </div>

            <AccordionContent>
              <div className="space-y-2">
                {data.subtasks.map((el) => (
                  <div
                    key={el.id}
                    className="flex items-center justify-between w-full p-1 rounded bg-slate-50 pl-8"
                  >
                    <div className="flex flex-col items-start">
                      <div className="flex flex-col">
                        <span
                          className={`font-bold ${
                            el.completed_at && "line-through"
                          }`}
                        >
                          {el.title}
                        </span>
                        <span className="opacity-60 text-[9px]">
                          {el.deadline
                            ? `deadline : ${new Date(
                                el.deadline
                              ).toDateString()}`
                            : "No deadline"}
                        </span>
                      </div>
                      <div>
                        {el.deadline &&
                          new Date(data.deadline) < new Date() && (
                            <span className="text-[10px] mt-1 px-2 bg-red-300 rounded-md p-1">
                              Overdue
                            </span>
                          )}
                      </div>
                    </div>
                    <TodoAction data={el} isSubTask={true} />
                  </div>
                ))}
              </div>
            </AccordionContent>
          </AccordionItem>
        </Accordion>
      ) : (
        <div className="flex items-center justify-between">
          <div className="flex flex-col justify-start">
            <div className="flex flex-col">
              <span className="font-bold">{data.title}</span>
              <span className="opacity-60 text-[9px]">
                {data.deadline
                  ? `deadline : ${new Date(data.deadline).toDateString()}`
                  : "No deadline"}
              </span>
            </div>
            <div>
              {data.deadline && new Date(data.deadline) < new Date() && (
                <span className="text-[10px] mt-1 px-2 bg-red-300 rounded-md p-1">
                  Overdue
                </span>
              )}
            </div>
          </div>
          <TodoAction data={data} isSubTask={false} />
        </div>
      )}
    </div>
  );
}

"use client";

import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import Todo from "./Todo";
import { AppContext } from "@/context/appContext";
import { useContext, useEffect } from "react";

export default function TodoTabs({ datas }) {
  const { state, dispatch } = useContext(AppContext);

  useEffect(() => {
    dispatch({ type: "INIT_TASK", payload: datas });
    return () => {};
  }, []);

  let ongoingData = state.filter((data) => data?.completed_at == null);
  let completedData = state.filter((data) => data?.completed_at != null);
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
      <TabsContent value="ongoing">
        <div className="space-y-2">
          {ongoingData.length ? (
            ongoingData.map((data) => <Todo key={data.id} data={data} />)
          ) : (
            <div className="grid place-items-center text-sm font-bold opacity-55 pt-6">
              {"There's no ongoing task"}
            </div>
          )}
        </div>
      </TabsContent>
      <TabsContent value="completed">
        <div className="space-y-2">
          {completedData.length ? (
            completedData.map((data) => <Todo key={data.id} data={data} />)
          ) : (
            <div className="grid place-items-center text-sm font-bold opacity-55 pt-6">
              {"There's no completed task"}
            </div>
          )}
        </div>
      </TabsContent>
    </Tabs>
  );
}

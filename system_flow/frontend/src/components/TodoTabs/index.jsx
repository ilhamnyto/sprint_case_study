"use client";

import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import Todo from "./Todo";

export default function TodoTabs({ datas }) {
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
          {datas
            .filter((data) => data.completed_at == null)
            .map((data) => (
              <Todo key={data.id} data={data} />
            ))}
        </div>
      </TabsContent>
      <TabsContent value="completed">
        <div className="space-y-2">
          {datas
            .filter((data) => data.completed_at != null)
            .map((data) => (
              <Todo key={data.id} data={data} />
            ))}
        </div>
      </TabsContent>
    </Tabs>
  );
}

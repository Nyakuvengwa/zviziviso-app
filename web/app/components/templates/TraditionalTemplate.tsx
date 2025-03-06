"use client";

import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Separator } from "@/components/ui/separator";
import { Cross, Flower } from "lucide-react";
import Image from "next/image";
import { DeathNotice } from "../../types";

export default function TraditionalTemplate({ notice }: { notice: DeathNotice }) {
  return (
    <Card className="border-2 border-gray-200 shadow-lg bg-[#f9f6f1]">
      <CardHeader className="text-center space-y-4 pb-2 bg-gray-900 text-white">
        <div className="flex justify-center">
          <Cross className="w-8 h-8" />
        </div>
        <h1 className="text-3xl font-serif">With Deep Sorrow</h1>
      </CardHeader>
      <CardContent className="space-y-6 p-8">
        <div className="flex flex-col md:flex-row gap-8 items-center">
          {notice.imageUrl && (
            <div className="relative w-48 h-48 rounded-lg overflow-hidden border-4 border-gray-300">
              <Image
                src={notice.imageUrl}
                alt={notice.name}
                fill
                className="object-cover"
                priority
              />
            </div>
          )}
          <div className="flex-1 text-center">
            <h2 className="text-3xl font-semibold text-gray-900 mb-3">{notice.name}</h2>
            <p className="text-xl text-gray-700 mb-2">
              {new Date(notice.dateOfDeath).toLocaleDateString('en-GB', {
                day: 'numeric',
                month: 'long',
                year: 'numeric'
              })}
            </p>
            <p className="text-gray-600">Age: {notice.age} years</p>
            <p className="text-gray-600 mt-2">{notice.locationOfDeath}</p>
          </div>
        </div>

        <Separator className="my-8" />

        <ScrollArea className="h-[400px] pr-4">
          <div className="space-y-8">
            <section>
              <h3 className="font-semibold text-2xl mb-4 text-gray-800 text-center">Family</h3>
              <p className="text-gray-700 leading-relaxed text-center">
                {notice.familyInfo}
              </p>
            </section>

            <section>
              <h3 className="font-semibold text-2xl mb-4 text-gray-800 text-center">
                Funeral Arrangements
              </h3>
              <div className="space-y-4 text-gray-700 text-center">
                <p>
                  <strong>Wake:</strong><br />
                  {notice.funeralArrangements.wake}
                </p>
                <p>
                  <strong>Service:</strong><br />
                  {notice.funeralArrangements.service}
                </p>
                <p>
                  <strong>Burial:</strong><br />
                  {notice.funeralArrangements.burial}
                </p>
              </div>
            </section>

            {notice.tributes && (
              <section>
                <h3 className="font-semibold text-2xl mb-4 text-gray-800 text-center">
                  Remembrance
                </h3>
                <p className="text-gray-700 leading-relaxed text-center italic">
                  {notice.tributes}
                </p>
              </section>
            )}

            <section className="text-center text-gray-700 pt-4">
              <p className="text-xl italic">&quot;Zorora murugare&quot;</p>
              <p className="text-sm mt-2">(Rest in peace)</p>
            </section>
          </div>
        </ScrollArea>
      </CardContent>
    </Card>
  );
}
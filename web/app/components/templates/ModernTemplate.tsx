"use client";

import { Card, CardContent } from "@/components/ui/card";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Separator } from "@/components/ui/separator";
import { Flower, Calendar, MapPin } from "lucide-react";
import Image from "next/image";
import { DeathNotice } from "../../types";

export default function ModernTemplate({ notice }: { notice: DeathNotice }) {
  return (
    <Card className="border-0 shadow-2xl bg-white overflow-hidden">
      {notice.imageUrl && (
        <div className="relative w-full h-[300px]">
          <div className="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent z-10" />
          <Image
            src={notice.imageUrl}
            alt={notice.name}
            fill
            className="object-cover"
            priority
          />
          <div className="absolute bottom-0 left-0 right-0 p-8 z-20 text-white">
            <h1 className="text-4xl font-light mb-2">{notice.name}</h1>
            <div className="flex items-center gap-4 text-white/80">
              <div className="flex items-center gap-2">
                <Calendar className="w-4 h-4" />
                <span>
                  {new Date(notice.dateOfDeath).toLocaleDateString('en-GB', {
                    day: 'numeric',
                    month: 'long',
                    year: 'numeric'
                  })}
                </span>
              </div>
              <div className="flex items-center gap-2">
                <MapPin className="w-4 h-4" />
                <span>{notice.locationOfDeath}</span>
              </div>
            </div>
          </div>
        </div>
      )}

      <CardContent className="p-8">
        {!notice.imageUrl && (
          <div className="text-center mb-8">
            <Flower className="w-12 h-12 mx-auto mb-4 text-rose-500" />
            <h1 className="text-4xl font-light mb-2">{notice.name}</h1>
            <div className="flex items-center justify-center gap-4 text-gray-500">
              <div className="flex items-center gap-2">
                <Calendar className="w-4 h-4" />
                <span>
                  {new Date(notice.dateOfDeath).toLocaleDateString('en-GB', {
                    day: 'numeric',
                    month: 'long',
                    year: 'numeric'
                  })}
                </span>
              </div>
              <div className="flex items-center gap-2">
                <MapPin className="w-4 h-4" />
                <span>{notice.locationOfDeath}</span>
              </div>
            </div>
          </div>
        )}

        <ScrollArea className="h-[500px] pr-4">
          <div className="space-y-8">
            <section>
              <p className="text-xl text-gray-600 leading-relaxed">
                {notice.familyInfo}
              </p>
            </section>

            <Separator />

            <section className="bg-gray-50 p-6 rounded-xl space-y-4">
              <h3 className="font-medium text-xl text-gray-900">Funeral Arrangements</h3>
              <div className="space-y-4">
                <div className="bg-white p-4 rounded-lg">
                  <h4 className="font-medium mb-2">Wake</h4>
                  <p className="text-gray-600">{notice.funeralArrangements.wake}</p>
                </div>
                <div className="bg-white p-4 rounded-lg">
                  <h4 className="font-medium mb-2">Service</h4>
                  <p className="text-gray-600">{notice.funeralArrangements.service}</p>
                </div>
                <div className="bg-white p-4 rounded-lg">
                  <h4 className="font-medium mb-2">Burial</h4>
                  <p className="text-gray-600">{notice.funeralArrangements.burial}</p>
                </div>
              </div>
            </section>

            {notice.religion && (
              <section>
                <h3 className="font-medium text-xl text-gray-900 mb-3">Faith</h3>
                <p className="text-gray-600">
                  A cherished member of the {notice.religion} faith community
                </p>
              </section>
            )}

            {notice.tributes && (
              <section className="bg-rose-50 p-6 rounded-xl">
                <h3 className="font-medium text-xl text-gray-900 mb-3">Remembering</h3>
                <p className="text-gray-600 italic">
                  {notice.tributes}
                </p>
              </section>
            )}

            <section className="text-center pt-4">
              <p className="text-2xl font-light text-gray-900">&quot;Zorora murugare&quot;</p>
              <p className="text-gray-500 mt-1">(Rest in eternal peace)</p>
            </section>
          </div>
        </ScrollArea>
      </CardContent>
    </Card>
  );
}
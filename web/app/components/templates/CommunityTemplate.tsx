"use client";

import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Separator } from "@/components/ui/separator";
import { Users, Heart } from "lucide-react";
import Image from "next/image";
import { DeathNotice } from "../../types";

export default function CommunityTemplate({ notice }: { notice: DeathNotice }) {
  return (
    <Card className="border-2 border-gray-200 shadow-lg bg-gradient-to-b from-amber-50 to-white">
      <CardHeader className="text-center space-y-4 pb-2">
        <div className="flex justify-center">
          <Users className="text-amber-700 w-12 h-12" />
        </div>
        <h1 className="text-3xl font-serif text-amber-900">Celebrating a Life Well Lived</h1>
      </CardHeader>
      <CardContent className="space-y-6 p-8">
        <div className="flex flex-col md:flex-row gap-8 items-center">
          {notice.imageUrl && (
            <div className="relative w-56 h-56 rounded-full overflow-hidden border-4 border-amber-200">
              <Image
                src={notice.imageUrl}
                alt={notice.name}
                fill
                className="object-cover"
                priority
              />
            </div>
          )}
          <div className="flex-1 text-center md:text-left">
            <h2 className="text-3xl font-semibold text-amber-900 mb-3">{notice.name}</h2>
            <div className="flex items-center justify-center md:justify-start gap-2 text-rose-600 mb-2">
              <Heart className="w-5 h-5" />
              <span>Peacefully departed in {notice.locationOfDeath}</span>
            </div>
            <p className="text-gray-600">
              {new Date(notice.dateOfDeath).toLocaleDateString('en-GB', {
                day: 'numeric',
                month: 'long',
                year: 'numeric'
              })}
            </p>
            <p className="text-gray-600">Age: {notice.age} years</p>
          </div>
        </div>

        <Separator className="my-8" />

        <ScrollArea className="h-[400px] pr-4">
          <div className="space-y-8">
            <section>
              <h3 className="font-semibold text-2xl mb-4 text-amber-900">Our Beloved</h3>
              <p className="text-gray-700 leading-relaxed">
                {notice.familyInfo}
              </p>
            </section>

            {notice.religion && (
              <section>
                <h3 className="font-semibold text-2xl mb-4 text-amber-900">Faith & Community</h3>
                <p className="text-gray-700">
                  A devoted member of the {notice.religion} community
                </p>
              </section>
            )}

            <section>
              <h3 className="font-semibold text-2xl mb-4 text-amber-900">Coming Together</h3>
              <div className="space-y-4 text-gray-700 bg-amber-50 p-6 rounded-lg">
                <p>
                  <strong>Wake:</strong><br />
                  {notice.funeralArrangements.wake}
                </p>
                <p>
                  <strong>Service:</strong><br />
                  {notice.funeralArrangements.service}
                </p>
                <p>
                  <strong>Final Resting:</strong><br />
                  {notice.funeralArrangements.burial}
                </p>
              </div>
            </section>

            {notice.tributes && (
              <section>
                <h3 className="font-semibold text-2xl mb-4 text-amber-900">Community Remembers</h3>
                <div className="bg-white p-6 rounded-lg shadow-inner">
                  <p className="text-gray-700 leading-relaxed italic">
                    {notice.tributes}
                  </p>
                </div>
              </section>
            )}

            <section className="text-center pt-4">
              <p className="text-xl text-amber-900 italic">&quot;Mweya wake urare murugare&quot;</p>
              <p className="text-sm mt-2 text-gray-600">(May their spirit rest in peace)</p>
            </section>
          </div>
        </ScrollArea>
      </CardContent>
    </Card>
  );
}
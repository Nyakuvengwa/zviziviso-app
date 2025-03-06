"use client";

import { Card } from "@/components/ui/card";
import { formatDistanceToNow } from "date-fns";
import { PublishedNotice } from "../types";

interface NoticeListProps {
  notices: PublishedNotice[];
  onNoticeClick: (notice: PublishedNotice) => void;
}

export default function NoticeList({ notices, onNoticeClick }: NoticeListProps) {
  return (
    <div className="space-y-6">
      <h2 className="text-2xl font-serif text-gray-900">Recent Notices</h2>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {notices.map((notice) => (
          <Card
            key={notice.id}
            className="p-6 cursor-pointer hover:shadow-lg transition-shadow"
            onClick={() => onNoticeClick(notice)}
          >
            <h3 className="font-semibold text-lg mb-2">{notice.notice.name}</h3>
            <p className="text-gray-600 text-sm mb-4">
              {new Date(notice.notice.dateOfDeath).toLocaleDateString('en-GB', {
                day: 'numeric',
                month: 'long',
                year: 'numeric'
              })}
            </p>
            <div className="flex justify-between items-end">
              <span className="text-sm text-gray-500 capitalize">
                {notice.template} template
              </span>
              <span className="text-xs text-gray-400">
                Published {formatDistanceToNow(new Date(notice.publishedAt))} ago
              </span>
            </div>
          </Card>
        ))}
      </div>
    </div>
  );
}
"use client";

import { useState } from "react";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";
import { DeathNotice, TemplateType } from "../types";

interface NoticeFormProps {
  onSubmit: (data: DeathNotice, template: TemplateType) => void;
}

const sampleData: DeathNotice = {
  name: "Sekuru Robert Tinashe Ndlovu",
  age: 92,
  dateOfDeath: "2024-11-10",
  locationOfDeath: "Bulawayo",
  familyInfo: "Father of Thandiwe, Themba, and Nomsa. Beloved Sekuru to many grandchildren.",
  funeralArrangements: {
    wake: "45 Matopos Road, Bulawayo, from 2024-11-12",
    service: "St. Luke's Anglican Church, Bulawayo, 2024-11-14 at 11:00 AM",
    burial: "West Park Cemetery",
  },
  religion: "Anglican",
  tributes: "Sekuru Robert was a pillar of our community, always ready with a kind word and wise advice.",
  imageUrl: "https://images.unsplash.com/photo-1556889882-73ea40694a98?q=80&w=2000&auto=format&fit=crop",
};

export default function NoticeForm({ onSubmit }: NoticeFormProps) {
  const [formData, setFormData] = useState<DeathNotice>({
    name: "",
    age: 0,
    dateOfDeath: "",
    locationOfDeath: "",
    familyInfo: "",
    funeralArrangements: {
      wake: "",
      service: "",
      burial: "",
    },
    religion: "",
    tributes: "",
  });
  const [selectedTemplate, setSelectedTemplate] = useState<TemplateType>("traditional");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit(formData, selectedTemplate);
  };

  const fillSampleData = () => {
    setFormData(sampleData);
  };

  return (
    <Card className="w-full max-w-2xl mx-auto">
      <CardHeader>
        <h2 className="text-2xl font-semibold text-center">Create Death Notice</h2>
      </CardHeader>
      <CardContent>
        {process.env.NODE_ENV === 'development' && (
          <Button
            type="button"
            variant="outline"
            onClick={fillSampleData}
            className="mb-6 w-full bg-yellow-50 hover:bg-yellow-100 border-yellow-200"
          >
            Fill with Sample Data (Dev Only)
          </Button>
        )}
        <form onSubmit={handleSubmit} className="space-y-6">
          <div className="space-y-4">
            <div>
              <Label htmlFor="name">Full Name (with titles)</Label>
              <Input
                id="name"
                value={formData.name}
                onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                placeholder="e.g., Sekuru Robert Tinashe Ndlovu"
                required
              />
            </div>

            <div className="grid grid-cols-2 gap-4">
              <div>
                <Label htmlFor="age">Age</Label>
                <Input
                  id="age"
                  type="number"
                  value={formData.age || ""}
                  onChange={(e) => setFormData({ ...formData, age: parseInt(e.target.value) })}
                  required
                />
              </div>
              <div>
                <Label htmlFor="dateOfDeath">Date of Death</Label>
                <Input
                  id="dateOfDeath"
                  type="date"
                  value={formData.dateOfDeath}
                  onChange={(e) => setFormData({ ...formData, dateOfDeath: e.target.value })}
                  required
                />
              </div>
            </div>

            <div>
              <Label htmlFor="locationOfDeath">Location of Death</Label>
              <Input
                id="locationOfDeath"
                value={formData.locationOfDeath}
                onChange={(e) => setFormData({ ...formData, locationOfDeath: e.target.value })}
                placeholder="e.g., Bulawayo"
                required
              />
            </div>

            <div>
              <Label htmlFor="familyInfo">Family Information</Label>
              <Textarea
                id="familyInfo"
                value={formData.familyInfo}
                onChange={(e) => setFormData({ ...formData, familyInfo: e.target.value })}
                placeholder="List family members and relationships"
                required
              />
            </div>

            <div className="space-y-2">
              <Label>Funeral Arrangements</Label>
              <Input
                placeholder="Wake details"
                value={formData.funeralArrangements.wake}
                onChange={(e) => setFormData({
                  ...formData,
                  funeralArrangements: { ...formData.funeralArrangements, wake: e.target.value }
                })}
                required
              />
              <Input
                placeholder="Service details"
                value={formData.funeralArrangements.service}
                onChange={(e) => setFormData({
                  ...formData,
                  funeralArrangements: { ...formData.funeralArrangements, service: e.target.value }
                })}
                required
              />
              <Input
                placeholder="Burial details"
                value={formData.funeralArrangements.burial}
                onChange={(e) => setFormData({
                  ...formData,
                  funeralArrangements: { ...formData.funeralArrangements, burial: e.target.value }
                })}
                required
              />
            </div>

            <div>
              <Label htmlFor="religion">Religious Affiliation</Label>
              <Input
                id="religion"
                value={formData.religion}
                onChange={(e) => setFormData({ ...formData, religion: e.target.value })}
                placeholder="e.g., Anglican"
              />
            </div>

            <div>
              <Label htmlFor="tributes">Tributes/Memories</Label>
              <Textarea
                id="tributes"
                value={formData.tributes}
                onChange={(e) => setFormData({ ...formData, tributes: e.target.value })}
                placeholder="Share memories and tributes"
              />
            </div>

            <div>
              <Label htmlFor="imageUrl">Image URL (optional)</Label>
              <Input
                id="imageUrl"
                type="url"
                value={formData.imageUrl || ""}
                onChange={(e) => setFormData({ ...formData, imageUrl: e.target.value })}
                placeholder="https://example.com/image.jpg"
              />
            </div>

            <div>
              <Label>Select Template</Label>
              <div className="grid grid-cols-3 gap-4 mt-2">
                {["traditional", "community", "modern"].map((template) => (
                  <Button
                    key={template}
                    type="button"
                    variant={selectedTemplate === template ? "default" : "outline"}
                    onClick={() => setSelectedTemplate(template as TemplateType)}
                    className="w-full capitalize"
                  >
                    {template}
                  </Button>
                ))}
              </div>
            </div>
          </div>

          <Button type="submit" className="w-full">Generate Notice</Button>
        </form>
      </CardContent>
    </Card>
  );
}
require 'prawn'

def background_color(color)
  tmp_color = fill_color
  canvas do
    fill_color color
    fill_rectangle [bounds.left, bounds.top], bounds.right, bounds.top
  end
  fill_color tmp_color
end

Prawn::Document.generate("print.pdf") do
  background_color "FFFFFF"
  #image "/Users/aa/black.png", at: [100,750]
  #image "/tmp/qr.png", at: [125,350]

=begin
  StreetEmployerBuildingTheorySidePackageLiving
  StudentDoorDebateApplicationFruitStudyPlan
  StudyMidnightAmountMinimumMessageMoneyMemory
  SuitSentenceStrengthEmotionBeltHandPower
  SunTechnologyLossSportMetalSkyCapitalDocument
  SympathyWingCriticismAgePhoneChocolateProgress
  TableEndHeartCaseCraftBonusQuarterMagazine
  TeacherProgressDimensionEggBlankIntroduction
  TrackSelectionHalfDotPointPhilosophyTrade
  TrafficEntryStrengthPaintMuscleMatterLoss
  UncleMannerLossConceptPerspectivePackClimate
=end
  stroke_color 'CECECE'
  fill_color "000000"
  dash(5, space: 2, phase: 1)
  line_width 1
  #stroke_horizontal_rule
  #stroke_vertical_rule
  stroke_horizontal_line(0, 100,  at: 0) 
  stroke_vertical_line(0, 200, at: 0) 
  draw_text "UncleMannerLossConceptPerspectivePackClimate", at: [0, 650], size: 12
  draw_text "TrafficEntryStrengthPaintMuscleMatterLoss", at: [315, 650], size: 12
  draw_text "StreetEmployerBuildingTheorySidePackageLiving", at: [0, 500], size: 12
  draw_text "StudentDoorDebateApplicationFruitStudyPlan", at: [315, 500], size: 12
end

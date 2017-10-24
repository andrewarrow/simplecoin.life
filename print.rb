require 'prawn'

def background_color(color)
  tmp_color = fill_color
  canvas do
    fill_color color
    fill_rectangle [bounds.left, bounds.top], bounds.right, bounds.top
  end
  fill_color tmp_color
end

list = [['StreetEmployerBuildingTheorySidePackageLiving', 'StudentDoorDebateApplicationFruitStudyPlan'],
        ['StudyMidnightAmountMinimumMessageMoneyMemory', 'SuitSentenceStrengthEmotionBeltHandPower'],
        ['SunTechnologyLossSportMetalSkyCapitalDocument', 'SympathyWingCriticismAgePhoneChocolateProgress'],
        ['TableEndHeartCaseCraftBonusQuarterMagazine', 'TeacherProgressDimensionEggBlankIntroduction'],
        ['TableEndHeartCaseCraftBonusQuarterMagazine', 'TeacherProgressDimensionEggBlankIntroduction'],
        ['TrackSelectionHalfDotPointPhilosophyTrade', 'TrafficEntryStrengthPaintMuscleMatterLoss']]

Prawn::Document.generate("print.pdf", margin: 0) do
  background_color "FFFFFF"
  #image "/Users/aa/black.png", at: [100,750]
  #image "/tmp/qr.png", at: [125,350]

  stroke_color 'CECECE'
  fill_color "000000"
  dash(5, space: 2, phase: 1)
  line_width 1
  #stroke_horizontal_rule
  #stroke_vertical_rule
  stroke_horizontal_line(0, 800,  at: 0) 
  stroke_vertical_line(0, 800, at: 0) 
  stroke_vertical_line(0, 800, at: 306) 

  y = 720
  list.each do |f, s|
    draw_text f, at: [10, y], size: 11
    fill_color '757575'
    draw_text "claim these coins at simplecoin.life", at: [50, y-40], size: 10
    fill_color '000000'
    draw_text s, at: [325, y], size: 11
    fill_color '757575'
    draw_text "claim these coins at simplecoin.life", at: [400, y-40], size: 10
    fill_color '000000'
    stroke_horizontal_line(0, 800,  at: y-80) 
    y -= 130

    break if y < 0
  end
end

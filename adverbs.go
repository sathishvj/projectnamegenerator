package projectnamegenerator

//from: http://www.momswhothink.com/reading/list-of-adverbs.html

var Adverbs = []string{
	"Abnormally",
	"Absentmindedly",
	"Accidentally",
	"Acidly",
	"Actually",
	"Adventurously",
	"Afterwards",
	"Almost",
	"Always",
	"Angrily",
	"Annually",
	"Anxiously",
	"Arrogantly",
	"Awkwardly",
	"Badly",
	"Bashfully",
	"Beautifully",
	"Bitterly",
	"Bleakly",
	"Blindly",
	"Blissfully",
	"Boastfully",
	"Boldly",
	"Bravely",
	"Briefly",
	"Brightly",
	"Briskly",
	"Broadly",
	"Busily",
	"Calmly",
	"Carefully",
	"Carelessly",
	"Cautiously",
	"Certainly",
	"Cheerfully",
	"Clearly",
	"Cleverly",
	"Closely",
	"Coaxingly",
	"Colorfully",
	"Commonly",
	"Continually",
	"Coolly",
	"Correctly",
	"Courageously",
	"Crossly",
	"Cruelly",
	"Curiously",
	"Daily",
	"Daintily",
	"Dearly",
	"Deceivingly",
	"Delightfully",
	"Deeply",
	"Defiantly",
	"Deliberately",
	"Delightfully",
	"Diligently",
	"Dimly",
	"Doubtfully",
	"Dreamily",
	"Easily",
	"Elegantly",
	"Energetically",
	"Enormously",
	"Enthusiastically",
	"Equally",
	"Especially",
	"Even",
	"Evenly",
	"Eventually",
	"Exactly",
	"Excitedly",
	"Extremely",
	"Fairly",
	"Faithfully",
	"Famously",
	"Far",
	"Fast",
	"Fatally",
	"Ferociously",
	"Fervently",
	"Fiercely",
	"Fondly",
	"Foolishly",
	"Fortunately",
	"Frankly",
	"Frantically",
	"Freely",
	"Frenetically",
	"Frightfully",
	"Fully",
	"Furiously",
	"Generally",
	"Generously",
	"Gently",
	"Gladly",
	"Gleefully",
	"Gracefully",
	"Gratefully",
	"Greatly",
	"Greedily",
	"Happily",
	"Hastily",
	"Healthily",
	"Heavily",
	"Helpfully",
	"Helplessly",
	"Highly",
	"Honestly",
	"Hopelessly",
	"Hourly",
	"Hungrily",
	"Immediately",
	"Innocently",
	"Inquisitively",
	"Instantly",
	"Intensely",
	"Intently",
	"Interestingly",
	"Inwardly",
	"Irritably",
	"Jaggedly",
	"Jealously",
	"Joshingly",
	"Joyfully",
	"Joyously",
	"Jovially",
	"Jubilantly",
	"Judgementally",
	"Justly",
	"Keenly",
	"Kiddingly",
	"Kindheartedly",
	"Kindly",
	"Kissingly",
	"Knavishly",
	"Knottily",
	"Knowingly",
	"Knowledgeably",
	"Kookily",
	"Lazily",
	"Less",
	"Lightly",
	"Likely",
	"Limply",
	"Lively",
	"Loftily",
	"Longingly",
	"Loosely",
	"Lovingly",
	"Loudly",
	"Loyally",
	"Madly",
	"Majestically",
	"Meaningfully",
	"Mechanically",
	"Merrily",
	"Miserably",
	"Mockingly",
	"Monthly",
	"More",
	"Mortally",
	"Mostly",
	"Mysteriously",
	"Naturally",
	"Nearly",
	"Neatly",
	"Needily",
	"Nervously",
	"Never",
	"Nicely",
	"Noisily",
	"Obediently",
	"Obnoxiously",
	"Oddly",
	"Offensively",
	"Officially",
	"Often",
	"Only",
	"Openly",
	"Optimistically",
	"Overconfidently",
	"Owlishly",
	"Painfully",
	"Partially",
	"Patiently",
	"Perfectly",
	"Physically",
	"Playfully",
	"Politely",
	"Poorly",
	"Positively",
	"Potentially",
	"Powerfully",
	"Promptly",
	"Properly",
	"Punctually",
	"Quaintly",
	"Quarrelsomely",
	"Queasily",
	"Queerly",
	"Questionably",
	"Questioningly",
	"Quicker",
	"Quickly",
	"Quietly",
	"Quirkily",
	"Quizzically",
	"Rapidly",
	"Rarely",
	"Readily",
	"Really",
	"Reassuringly",
	"Recklessly",
	"Regularly",
	"Reluctantly",
	"Repeatedly",
	"Reproachfully",
	"Restfully",
	"Righteously",
	"Rightfully",
	"Rigidly",
	"Roughly",
	"Rudely",
	"Sadly",
	"Safely",
	"Scarcely",
	"Scarily",
	"Searchingly",
	"Sedately",
	"Seemingly",
	"Seldom",
	"Selfishly",
	"Separately",
	"Seriously",
	"Shakily",
	"Sharply",
	"Sheepishly",
	"Shrilly",
	"Shyly",
	"Silently",
	"Sleepily",
	"Slowly",
	"Smoothly",
	"Softly",
	"Solemnly",
	"Solidly",
	"Sometimes",
	"Soon",
	"Speedily",
	"Stealthily",
	"Sternly",
	"Strictly",
	"Successfully",
	"Suddenly",
	"Surprisingly",
	"Suspiciously",
	"Sweetly",
	"Swiftly",
	"Sympathetically",
	"Tenderly",
	"Tensely",
	"Terribly",
	"Thankfully",
	"Thoroughly",
	"Thoughtfully",
	"Tightly",
	"Tomorrow",
	"Too",
	"Tremendously",
	"Triumphantly",
	"Truly",
	"Truthfully",
	"Ultimately",
	"Unabashedly",
	"Unaccountably",
	"Unbearably",
	"Unethically",
	"Unexpectedly",
	"Unfortunately",
	"Unimpressively",
	"Unnaturally",
	"Unnecessarily",
	"Utterly",
	"Upbeat",
	"Upliftingly",
	"Upright",
	"Upside-down",
	"Upward",
	"Upwardly",
	"Urgently",
	"Usefully",
	"Uselessly",
	"Usually",
	"Utterly",
}

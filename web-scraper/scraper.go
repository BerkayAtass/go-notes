package main

import (
	"flag"
	"fmt"

	"github.com/gocolly/colly"
	"github.com/mbndr/figlet4go"
)

func init() {

	logo("Yavuzlar", "Web Scraper")

}
func main() {
	var targetItemCount int
	var desc, date, price, one, two, three, help bool

	flag.IntVar(&targetItemCount, "d", 10, "Number of items to scrape")
	flag.BoolVar(&desc, "desc", false, "Include description")
	flag.BoolVar(&date, "date", false, "Include date")
	flag.BoolVar(&price, "price", false, "Include price")
	flag.BoolVar(&one, "1", false, "Scrape HackerNews")
	flag.BoolVar(&two, "2", false, "Scrape Trendyol")
	flag.BoolVar(&three, "3", false, "Scrape Ebay")
	flag.BoolVar(&help, "h", false, "Help")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	if one {
		fmt.Print("\n\n\n\n")
		fmt.Println("Scraping HackerNews")
		news1(targetItemCount, desc, date)
		fmt.Print("-----------------------------------------\n\n")
	}
	if two {
		fmt.Print("\n\n\n\n")
		fmt.Println("Scraping Trendyol")
		news2(price)
		fmt.Print("-----------------------------------------\n\n")
	}
	if three {
		fmt.Print("\n\n\n\n")
		fmt.Println("Scraping Ebay")
		news3(price)
		fmt.Print("-----------------------------------------\n\n")
	}
}

func logo(yamlFile string, yamlFile2 string) {
	// YAVUZLAR
	yavuzlarASCII := figlet4go.NewAsciiRender()
	yavuzlarOptions := figlet4go.NewRenderOptions()
	yavuzlarOptions.FontName = "larry3d"
	yavuzlarOptions.FontColor = []figlet4go.Color{figlet4go.ColorBlue}
	yavuzlarStr, _ := yavuzlarASCII.RenderOpts(yamlFile, yavuzlarOptions)
	fmt.Println(yavuzlarStr)

	// Web Scraper
	webScraperASCII := figlet4go.NewAsciiRender()
	webScraperOptions := figlet4go.NewRenderOptions()
	webScraperOptions.FontName = "block"
	webScraperOptions.FontColor = []figlet4go.Color{figlet4go.ColorRed}
	webScraperStr, _ := webScraperASCII.RenderOpts(yamlFile2, webScraperOptions)
	fmt.Println(webScraperStr)
}
func news1(targetItemCount int, descT bool, dateT bool) {
	c := colly.NewCollector(
		colly.AllowedDomains("thehackernews.com"),
	)

	itemCount := 0

	c.OnHTML("div.clear.home-right", func(e *colly.HTMLElement) {

		var title string = e.ChildText("h2.home-title")
		var description string = e.ChildText("div.home-desc")
		var date string = e.ChildText("span.h-datetime")

		fmt.Printf("%d. News: \n%s\n", itemCount+1, title)
		if descT {
			fmt.Println("Description : \n" + description)
		}
		if dateT {
			fmt.Println("Date : \n" + date)
		}
		fmt.Println("-----------------------------------------")
		itemCount++
		if itemCount >= targetItemCount {
			c.Visit("about:")
			return
		}
	})

	c.OnHTML("a.blog-pager-older-link-mobile", func(e *colly.HTMLElement) {
		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
		if itemCount >= targetItemCount {
			c.Visit("about:")
			return
		} else {
			c.Visit(nextPage)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	if itemCount >= targetItemCount {
		c.Visit("about:blank")
		return
	} else {
		c.Visit("https://thehackernews.com/")
	}
}

func news2(price bool) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.trendyol.com"),
	)

	c.OnHTML("div.prdct-cntnr-wrppr", func(e *colly.HTMLElement) {
		e.ForEach("div.p-card-wrppr", func(_ int, h *colly.HTMLElement) {
			if price {
				var title string = h.ChildText("span.prdct-desc-cntnr-name")
				var price string = h.ChildText("div.prc-box-dscntd")
				fmt.Println(title + " --------> " + price)
			} else {
				var title string = h.ChildText("span.prdct-desc-cntnr-name")
				fmt.Println(title)
			}
		})

	})

	c.Visit("https://www.trendyol.com/erkek-t-shirt-x-g2-c73?pi=7")
}

func news3(price bool) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.ebay.com"),
	)

	c.OnHTML("div.item-grid-spoke", func(e *colly.HTMLElement) {
		e.ForEach("div.col", func(_ int, h *colly.HTMLElement) {
			if price {
				var title string = h.ChildText("a[itemprop=url]")
				var price string = h.ChildText("span[itemprop=price]")
				fmt.Println(title + " --------> " + price)
			} else {
				var title string = h.ChildText("a[itemprop=url]")
				fmt.Println(title)
			}
		})

	})

	c.Visit("https://www.ebay.com/globaldeals/home/toys")
}

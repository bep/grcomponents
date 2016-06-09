# GR Load Bar

[![GoDoc](https://godoc.org/github.com/bep/gr?status.svg)](https://godoc.org/github.com/bep/grcomponents/loadbar)

This is a load bar for [GR (Go React)](https://github.com/bep/gr) that can be used to indicate loading of some slowloading content in the browser, iframes etc.

## Use

```go
var loadBarIframe = loadbar.NewLoader()

// loadBarIframe.SetStatus must be used as a callback to report back status changes.

loaderElement :=  loadBarIframe.CreateElement(nil, mySlowLoadingComponent)

```

It will wrap your component(s) in a div with CSS classes.

Example CSS:

```css
 .gr-lb-wrapper {
       padding-top: 3px;
       width: 100%;
   }
   
   .gr-lb-wrapper.gr-lb-initial .gr-lb {
       visibility: hidden;
   }
   
   .gr-lb-wrapper.gr-lb-loading .gr-lb-component {
       opacity: 0.3;
   }
   
   .gr-lb-wrapper.gr-lb-loaded .gr-lb {
       visibility: hidden;
   }
   
   .gr-lb {
       display: block;
       position: relative;
       margin-top: 0px;
       z-index: 100;
       width: 100%;
       height: 6px;
       background-color: #fdba2c;
   }
   
   .gr-lb .gr-lb-bar {
       content: "";
       display: inline;
       position: absolute;
       width: 0;
       height: 100%;
       left: 50%;
       text-align: center;
   }
   
   .gr-lb .gr-lb-bar:nth-child(1) {
       background-color: #da4733;
       animation: gr-loading 3s linear infinite;
   }
   
   .gr-lb .gr-lb-bar:nth-child(2) {
       background-color: #3b78e7;
       animation: gr-lb-loading 3s linear 1s infinite;
   }
   
   .gr-lb .gr-lb-bar:nth-child(3) {
       background-color: #fdba2c;
       animation: gr-loading 3s linear 2s infinite;
   }
   
   @keyframes gr-lb-loading {
       from {
           left: 50%;
           width: 0;
           z-index: 100;
       }
       33.3333% {
           left: 0;
           width: 100%;
           z-index: 10;
       }
       to {
           left: 0;
           width: 100%;
       }
   }

```


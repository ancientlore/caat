package main

import (
	"github.com/golang/snappy"
	"encoding/base64"
	"log"
	"mime"
	"net/http"
	"path"
	"strings"
	"time"
)

const (
	c_CatTxt = "mdoCRBtbNDg7MjsxMDs5OzEwbRtbMy4PAAjiloQZIQQxMB0iAS4BDx0iBDM7AQMAMhVEBDE7BQNGIwAAM64jAAFmUmkAGUYYMzsxMjsxMvppAEqvAAF5AXwxFgkjMmkACRMAMREjEDY7Njs1HYkQNTs1OzYVHQQ5OwUDHSAAMgECESAEODgJA14gAAQ1NQUDADYdIAX5AfwRoymFADMuIwA6owAANwECHSAAOQECFUARLTnYNiAAAW0BcBkgARMBdhVDADJG-AERVhGmERAdRiFpLmwBAWY6aQABIC4jAAENADsuTgE2qwAFHwQxMB1lEDg7OTs4FYUJIDKgAgAwBbt5B4rLAAlGLmYAARMAMQF5Ee4RVmruAFZGAAEWLkYABRA2RgAhSm4jAAEgIUc2jAA-9QAFWTYjAD7SADojAAA5MnkBDWcZigl3OmcAATQRZUohABETESMpNy5EAH7gAQEzOqgACRA5NDJDAErFBD6FABEQGUIVExGoADKN-B0jAYVdAgl1GSAJEA1AYZAu8wM-YAAENDYJAx1ABDE1CQMRQwQ1OAkDHSMyYwAVkBmDFRMNhgkwZiAAFZa2IwAxPw1mFRAZRgl5DSAJDRkdFTANIBUQHSMQMzsxNDs2DgaFjx0jgUkuBQMBDQEQHSMBMy5GACnVrUgIMG0K7QcJoTaOAHoqAwW0LmsAQRpdHQ17MlYGAYsujgAENzsFAy5GAAA0BfcANBGMBDIwAQMAMS4OBUIjABQ0OzIzOzIN1On3EVkN0gUgOiMABDg7BQMVIwQ0OwUDHYxBwgA3ESAV3zldGDYzOzYyOzbZIgAzIV4FBB0mCDE3NBEEFUwIODU7DQQuKQAINDk7DQQVKQg3MDsNBC4pAAgzNjsNBBUpCDY5Ow0EHSkYODc7ODY7OBnECDI1Ow0EHSYlBwQ7MTUKAdQIMjs2MsQAhvACPiMASVI5LQkQLZAVLRkgQhwCEXOCIwAyqQCpXdleSuoGQaVdqAmZHWZhQ2FGHawBRi4jAEHFLsgCAQ0BEB1GAWMdZhUQGSMJeRFmDqkKLrwBPh4ABUEdPgUQHWFK9AUBHx0iBWOC1gMFMhmlBRAAO1mRAR86qAAFFB3ICRA-QwA9lAVESokAFWYB_DIgAAERFR5BGgFDRmMA9ahKDQEOuggAOxkiBUI2ZQA2Rgkh0HrTAWE1YTgd61G1MdYAN2VePuEEBVkNIwQyNgkDHUYENjQJAxEjADEBagkEHSYEMTShK8EmGSkIODY7DQQuKQAIMzg7DQQVUgg1MDsNBC4pABQ1OzE2OzGZVQA05Y8uEwUlWTWKADgBAi5BAA6TCS47A6EMOkUBASAuIwAANaXQMpAEPkUBERA5aDIQAgkNGR0BYwEzFe0BDSH7HSMBlgGZLYUVEDYjADKpAIEEAbkdIwEgARMRRgFmFWn5XSFIIawVKEGJNnwCFTsRSwEwADsyewElmzqeAT4jAAQ5OwUDFUYREBm0BDIy6Q0RI8H0CDM7Mi7kAeFzCDE7MjVAATMIMjsyLmcFBDI5CQMVRhEQGWkENDIFAxVpBDM4CQMdIz7SAAFmBXkZRgQxNEEgBQQRbAgxNDARBB0pGDIwNDsyMDMBCBUpCDk3Ow0ELikAAe8JBBEpADIhBQkELikACDE3Ow0EGSkINjsyCQQdKRgxOTE7MTkyAQgVKQgxNTsNBC4pAAAyQUEMMDsxMrkvADkhowUEHSkENDEJAxFPCDEzOREEHSZJ-E0xBDQ0CQMdID5UAkKuA6EAQYcNQxUQOX8yYwAVUF4gAAFNAVBeIABhHHUfFqgJYT8uQgMBDQEQGUN1mxGmShYRAXYBeREjAWY2aQAyyQAJDRljhdWKoAb-IwCGaAABmkJmAC56CAEhdXZS6AAlDTUKDlYIakoHITY2IAA1gjFMASB6IwBhnzqsBKFdvWBhcDqAC2HthVYNaWX6gXY5mGGTDpYLESMEMzEJAy4jAAFDgZwNRgAxAXoJBB0mBDEwDgoUBQQRKYkwgTQyMwMINzk7DQQVKQGDCQQdUhgyMDc7MjA4AQQVKQgwNjsNBB0pCDE0MREEESkIMTA4EQQdKQ5iE_l6CDsxOy7nADWGEechXSFgOS0BI8VVESNhAjq5AQ5QDgE2ESNhazpuAyHsATYRI4GyNiMAMoMCCQ0dhgl2GewREB0jAYMuhgABDQEQHSMRNh2GARMdIwl5DSAJDQVA_XYOGwkBWA0lBUI2FgQhAQHhFSMegBAFS4miPkoOEpAOADIyewZBpOFDDWkVEBlGBDI1CRYVIwQ3OwUDLiMA8UMNRgQzMgkDLiMACTZVck4jAAQ4OwUDFUYUNTszNDszMrsCADUFj_VQFRAZr1J2B1nIGSntSQEIEXIEMjMOYggFBC4pAAwzMDsyAeMAM5WcBDIzDrgPBQQyKQAcODsyMzc7MjMVxAgyMzkJBAA0MnYBCU4AOy5SAAg0MjsNBC5SAAHhCQQZpAg2OzIJBC4pAGXUZcUVKSFzCQQdKQQxM-HEDDk7MTKVJggxOTURBB0pEkcPDkoPLdQAMQ6RCQkEHSZVdg0mGDkwOzg5Ozku7QDpuw0gYcdKcAUAMmGnESMBDTrsAmFkLg8DqbYurwBhDA6bChVGATABMznxBXmh7BFpATCBMB0jgUABExEjDrUKADsyaQAAOQHRFSABYzZmAGnQEUBOhgwAOP0GabUuRQIByWXrHUMFth2mKSJ5-wEgAdk2IwAFmRFGAQ02qQAVExEjFvEMLn8CQiMACa8yeAEBeQF8DUZBcWHrLowAKUVVme0LLiQCgXflZBFGSqoE4WH9Lhg0OTs0ODs0LocGBDM2CQMRRgAxoQEJBC6PAAHGCQQVKRJ8Dg7wDjZUAwgxOzIJBBEpbYwBCB1SADKBtwkEFSlx-AAzLi4DhUqBTi5KBAH4QkoEAe4JBBVSEDEwOzIxIQUuSwYAMSFdCQQRKQQ2NgUDLvMAMtIa6a1dK0EYQRstfA6tD34bEEoIAgFDLkYAQcQ6xgMBEAETEWlhamFtGYw2PQMBLQEwGSAVExFDASA6IwABhi6JAIFMgV-yRgA2iQCJKS14_aY6YAci-wwda2EcLh8DIcE2KgqhVA5FDRHUTvwCDicKXfxKRgDBJeEdFUZhhQQ7Mi4GAuVN4VAtPQQzMAkDHa8EMzMJAxUj6ckN-hJVCOGW4ZkVIwA1CRMdRmXLDrIWDWkNEDJaBBJPDiGYVeEYOTM7OTI7OS78CggyMTgRBBVMCDE5Ow0EMnIACDE7MgkEFSkANA7YCA4fEC5eBS65BxGeEDQ2OzI0DuUIMv0D_ZARKQg1MDsJBC4QAQkpAWi1kAlOAVI2FgEhFAUEFVISLQgOMQguqAGNSwEIFSkOMQgJBB32CDE2OBEEFSmJsIG4LikAZQUIMzsx3TRBdGHoADEuQgUOdRcOeBcRTxomGDK6AYEcDjoUMRYe_AlZ9aHOIX0RIxUzHSM67AlFvAQ7Mi4wARU2DWm13hlGPoEDSvUCAed6XgOhP2HKEWkhLQAyOjIEifEV0okfHYmBcjp1BD5DAA7PDnXlSjYDDlkIFSAOqREOXRUdZG1GVT1KEAEFMwQ7MjW_7WouCwI1Iy1WIWYAOSHPGc2Bbi5qBwEQhX4ZI5XqDUYENTkJAx0jBDk3CQMRIxa8GQA2AQgdJggxNzURBBEpCDIwOQ0ELpoCCDIxNBEEFSkB7wkEHVIAMmVj6VENnmn1ADUBCDIpAGV5DrILFSkhHgkELikAIXoJBBUpnZk5FggxNDIRBBEpGDEwNTsxMDQBCC4pAEEcQWINoQAwAQJqIABBfw5FCS4jAEmSNTYVEB2PQZIulQLhKjYtBzacCwEdNiAAFRMNqRW5amYAAXMBdh0jAVYBWRFGAQ0BEBkjaVENIBJsEi5MAhUxDSHhkeGUHUEVdxEjIWAhYwVk-bQENDAJAxEoARAEMTsyng0i7hUNSwEjATYJS-UF9aQNI0q_BGFD5eoRIwENNvoHoSWF9REjAQ02BQW1XhEjADdp5R2M7aRVc07XB-H3LvoH6dcu1wU-kQUJRgAyLoYBGDUxOzUwOzU1O2ETDj8PHYwAMRJFCw5JC3VUaRUO0BAuKQAIMjQxDQQdT-UcYcEdUunHwU55-FZcC_WeGVIhrAkEMlIA4fAFBDVWHDUxOzI1MjsyNtoHiUYFkTF_VqQARm4XCDE3OBEEHXsIMTY3EQQRUgQxMEHuBQQdKQQ2MAkDESYVEEXOAdPR2fX6Tg8GIU0qiAkENDcJAx1GBDM3CQMRaQQ4MwkDHSMOaQ0IOTs01ZsIMTA2EQQdJiIqCC0qBTNBLRmyIjUjESMOvw12wg1JhjIrDmHZDkcXEUYiswoZaQEjZfwNacEkADXhFxkjbY41YTUIGSMiwSMNRi0-LtABQtIAISghKxlG9ekRRgAy5T0uDAMENjUJAxEjBDgxCQMuRgBxsQ1GBDcxCQMdIw3SIgALBDY4CQMuIwCBXYFgDUYUNTc7NTY7NrwoADAOkg0FBBEmSkkACDE5Ng0EFbUWUAoAOQEILnIAEkwa5boRTwgxODINBDK-AIGACQQVKRKbGg6fGjIpABb7Egg7MjMVewgyMTYasRoukQIEMTgOpgoFBBVSDvgSFgATHaQQNDA7MznBAjEWDiEqHSAANQ4yDxEdwXchHh0gQp0GQXs23gY-Bg5BwT4kBwk2VdcFZjpGAAEz4VoNqRpSCllsQiEA4WV6RAARdx1GATMBNhlnAY0uhQE6rgcORCgdQDLpCgE7OjEBge0O9w82YAAtU_mqBDM5CQMRKOGXga4da-WXPbvhygQxOzJdA2H8LhML4UE6GA8iOwkRkQ6XDiEBGdQ-XQMVEB0jwaQ6ShelfR0j8b0NaRUzHSM-GA0BvDa7BIEMDk0aEUZKAQU-nwhO5gFKiCVS2QUu8wYRuFaEDyEVCQQVKQgwNTsJBDJ5CwH4IRwmzAwuaCY5DQgxOTARBBFSJmEILtkETuwDUlUmGDE1MzsxNTIBCBVSCDYxOw0ELnsAhcsAOA6_IhUpCDA3Ow0EHSkEODkJAxEmoVQINzs1LiMARbYEOzP1ZBYUDA4YDC5nBAgxMDMRBBFMLqQBPRYSTwwOUwy5VBU8LlgGTikAFsIfEsYfGVLhvAAzLgwoUuoAYVoO3B8RngA04d827QYOsScINjs4uS4IMjg7DQQubwAWoRQZJl2LGZUENTMJAx1v5V8uIwAOjgoOZAhNmAQ3NwkDHSMENzUJAxEjLuAGHWwerCsNSUp3CAQ5MAkDESMEOTYJAx1GFmYRgTq1xiUOIREZbAQ4MgkDEUkEODUJAx0jPjsPBDgwCQMdIwQ3MgkDEUYJ2A7OGC7GAiInCA3YVuIDBDkxCQMVJkn4OssHADEWxAsO0AsVKVaTKQg0NjsNBBUpCDYzOwkEMqACKiQLEaEqIQs5MwAxFjcQAQgVKQ6yEQkEHSk29SoeIg8dScF43XsW_w4dICG9ofoRj0oeB-GCLoUHAQ0BEB1GAUMuRgABDQEQGSMyQQcNli65AQEgQkwhLgACCTSVduVTLhsDAUEVHgktMkAABTGKYgA-IwA2ZADh6zUsIVkBMAXm-ccO8QoOTyURKAQ0MxLtIzJiDAEg4bctdwEjBDM7MiMAPrsK1eIFbhIRDX7zABJCEQ7MGw1mTt0BDioUDrEtESMB0RXUBWbBvOXXDSNKiQASLRMqrhdhTw6ICh2sAc8yLQgOGhA2HRA-NApKIwCV8A2MSrcEBDE2DjgSBQQRJgQ3MA7yNAA3Mk4CCDkzOw0EESYWXBYOmS0yuABG6g4aSiIBCB3hGiIuAQgVUioeEzm27XsBCBUpxaUOhQkypAAqeiIN8C7hBhlSPpIHLjkAGSYu8C4RTwAzYckOdyEunQLpDRL2IhEpBDU0CQMdTz6QBQ5yEQ5MGR0jGnQcAQgRSQgxMzENBDLqAAgyNzsNBBEpBDk1CQMdTw6oFwA3KiAvIrAfHeS1dznXEhINDqERLq4BIuU2LQoFWQFcHUy97A0mSSMSgxgdKQQ3OAkDESbVXh1MYWgJBBUmADIS_RkyHwYOqg0FBFUaDiMaADM2IxoSwQ4OxA4NlRpVEC4WBinCEr0JESYIMTc2DQQuKQAO6g4INzs2FSYWggwyHBQ-mwcIMTM3EQQd4QQ0NRIbEDnowSZh0TbqHxYECg4QChGYEjsKNj4KSoUlNa85Uw64EQg0OzgiEgohwjp5CC79By0tADlKxQhl9GH4NZwAMcXBwcIuowYEMTUO3CAFBBFPBDc0CQMdmAQ3NgkDESMO6yk62wWhiAFdESPhxTJXCaGCoYUNtQUtNk8HAUPhqxUjATABMzkhIYJh3xUjATAiIBal2w0TciMAAUPh7hGMAUM6RgABnAGfESMOxg8BEBmMfqoG5Z05fwHKNkEACRMZIgFgNu8AfmMADTM1dRUQBab9vgA3Rco5wAA4FggTCSglNxIVDAQ7NBVL4c4AM-HOBUsFI8X_DmUaMRQObxsAOzZbMKHwIUoRIwnUMnICASAuIwAhigEQHWkRExFGwfrh9B0jCRMV0gEQDpYVLiMAoSAdr0rkBwEjAUkVaQl5MvAFASAB5Q2M4cEAMDp-JwEQARMRIw6eEw6hEx2vDj4gKkEgoTcSowgZI4G3Du0TEUYilhQZI419AQgRJoEiBDc7MgcGtasNjwQ3OQkDHUntXRVsGu8xLpkGSrEqSkkAEsIxDjU2Jq8XGDE4MzsxODQBBB1yzSkBCBGYBDE3QT0SYRcdKULtDAg0ODsNBB0mBDUyCQMRTAQ5NAkDHSMIMTcxDQQi7g0uFgU5Vj71BwQ5OAkDaiMADvwOCDc7Ni5uAhb6Lw7-L5V6GgoVDnYoHUxKgwcW9ycOzysukgMW4hAO5hAV5y5sDRnBFf0t8S62CR0mPsodUigSQiYAFvowDpApHUzhrw4yPxVyYTkFBC6CAS5UCg2bLkoWHZseNwgNJhYrNOWTHSY-qAGBnz6FBkq1DmHBCQQdTCL8CQ1yIs8PGZVhFAg1OzZ1allQNs4EDgwfBQQRTBbdEoF6LkMCFWIRckYfGBrAECK5CLWTGZUaiBgVlQ0QLpwBADMBAhFmyUEZQLWQDYYhtKWNGSMJMw0gCQ0ZHQFAxcYNIAENwdkuIwABEMHHZiMACVbmIAAFs8EzEUOKZgYBMAEzNswAATkVRgEgOiMA0RARaaHEoccdRgHsLu8AAVM26Adq7wD9ow5EGw5HGw1rIqcdGWsiHBwNI0r9AcHRwfcRI0H5wcEdRg5REQXEDUbVvhlpBTMAO_1dwasOUQ8dIwFD4ZMNRgEjQfMuIwBKPQEOzREuIwABZnppACFNIVARacE14ZAZjA6jDQgyOzNV9SE7BDE7MgUe7U15WyFrIW4ZRhIQECpnHQFTBDE7Ml4fAeUlCA2MDgoYOoYYNZQNIwFGDvYbHWkiCBQNIwQ3MAkDHSMSyBUqyxUENzMJAx0jIkwODUYENjEJAx0jBDE2Dk45BQQRJhrHFQEILikACDcyOw0EFSkWrDIOBwguKQCllRLMORUpCDY0Ow0ELikASrQOCDU5Ow0EHSkWWBgSYBgZUg6NDg68DDKkACpLNw3zCDE5OBEEHVIunTcNKQQyNQ6PCQUEHSkWhysSjysVKVYRJBJaNw6sHxUpCDE1Nw0ELnsA6fEANwEIFVIWzygO1yguewASXRIOXhI1cVaULAgxNDcRBBVSEo0dEt03LlIAADAOggoFBBUpVkUoJUghORVSVlYICT8AMgEIEVIWhxIOZws2UQQOnAsBBBVSGh4tAQgdpD5cCAAyDhcjCQQdJi49FjHASSFBLS4pAA4IHgkEESkWdBI-GRkIMTYwEQQVKRZlFwQ7MTJpF20EAQgZKaHdBQQuewAWShIOlRIRKQ67Dw52IR0mDj8ILkIIStAHEj8ULm8cgYIBEADiJh9KPkcFoQ6hER0jgZWhoC0wSiQFQqcG4cY6RgABEAETEUYB9sHaGWkyygYJDRkdDUA9Owl2LiMAHp0wEWPBQzpQBwAy6d8VIxH_HWkRExFGERBqIwAB2QHcaiMAKRIyrwARWRFpAUNGIggBshUjAQ0VI_nfPnEKFRAZkSL4FQ1uxTh2gQY1-RlGhVYqtxYVEBkjPnYHTgIIQU9dUg72DjraFk2r9XYVEB1p4anhrA3SSiMADmIVMtQmHv0xGUYSagwEOzXVIPVjHSMOQBwqQxzl_wA7NgIIDs0QDo4YDYwhgQ6cHB1GFikUADOVTg4DERLRGBlpPiMAIhUeGSMONhQENTsmNhQOfBEOsiUdIxpbDBXSLhMOGUkuoQ0NtS7SBx0pIowjVYcmmScudwYmMBUmrQwquisZUi4xBw17CDI0NxEELikAxe4MNzsyMrn8UkQHCSkOvD4ZKQ4oEQkEHVIIMTgxEQQVexLYJw7cJzLNAA5XIAA4DvUsNbHhBQ7PPgQyMC7GAwgyNTMRBBVSUvsgFrkNDj8LGaQSlg8OFxku9gAaZRkBCBFSLq0PAOK1vCYkISKyDVbMBhpJIQEIFVIiniUybQcq5jctmi5sCRl7LuwQESlSzA8AMuXk4egVe1aaPBqlLQ6qCBFSFiQIQqgwEmMIEvpAFSkWfgkIOzI1LiYKCVIOiDVVkFoWKhYUOQ7dPBVSEi8mDjMmLh8BLnkIDfZWMAlNOgEIESlWNxEEMTKhzwUEESku2BA5cbVGDXguKRoZJtVDDSZhEzoaJBIlDS5bDQ5YDQ7IFB1GASAuOA2hZlIjAFl4IX06lRQOew29n0LRB_HBDawWWhAZhgkwDR3N-TKiAxYkDhWD4U8OmBsdQwET4WU6RgAydAzpdZUOBTAOvRcZiRUTEYwBIPWFFrwWIcIBExEjFooIMowAPmkAERAZaTLyAAkNBR356BZPIyYOFU5pDg7xCQ7kCQ2LTqIHIfAh0A0jDRAuhAQBNuWSDSMBEOX4BZEF9D6SAgUgDhsIZiMA9fgdIxaTCjWjFRAdIwFWAVkNjO0WLpkD5eUqvC7lWTZfGPX7EUZGkzY-6AdKRgD1-x1GBXkZrz5wKAEgEgw3GSM-6AcaQBAu0gAarhMiAQwatA8yuwIW4CUOESYRj0pdGyI7JA2yFRAZjy5ZJw0mVnIuFoYWoc2Vri7BBBlSFgMgDpg3taRW2h8u3TYRe6ncDjwsLlIAEvEz4Su1Uun1ADcBCC4pAAg0OTsNBBFSGDIxMjsyMTMBBB0pADESVQ4OWQ619iYjJzIpACrFCA2kEksZADs6QhAe7ywZT1LaJi4oBw1PLicdAOL53xZxDxJwCBEp_aAZKQgxODARBBUpBDczDmgOAQQuKQAJ7wH7GSkmXDUdUhKCDjIOHyISEh0mKowZDaEOfAoO3UsuJgBFpxK4KBEmSpweFiI5Et9AESYWzTwOQhEycAmFD4EQGSkSjkGBSS7tAC5IBw2eCDE5NBEEHaEuZwgRKRLmIaHoMgkCEgI-ADYOgkwVKRaWMWEwHVIu7AINUiLkHjk_DicbDjlYESMSY1gEOzYuSQAOvxAEOTsmEUPhZzrQBzISOKHKDnYIHWYOfyUOZw8NiOEj4TkdIiI4Mg0jBWI2Tyw-HRkEOTkJAx1GrZZ18E7oGg63HxFmDjINPmAZ8RMNie1pMoMB4WYuaQcJMy74ARU2qkYAIRI6KRAh_Q4ICBFGDlsIOiMACXwmOwgREDnj6cUNQyGHIYoFH_nUPnYG1WMFKOlrDm5E_QLhJQETHSMh-jLUBxbECTL8DBbHCTWCGqEJMmYFYVgdaQ32LmkAEhcIDoYmDdahgTqmFlVBDSMiYw8Z0hUTDSNK_AIOeR4SYCQNIxKrIjYnKw7oIxIOPg0jVo0GdXsRJhJYUCHwLh4B9YQRJhaJFg4uHB24DpsiDolMFSZS4wUiURcRTBbDGxKcKjkEdUURJlL4TS5MFBEpADIOfwoFBGpPAIFcJUAZchZnTQHU9RxO1wUSHhcOIhcVJi4bIC5PAGEiBQQZmxIfFQ4rFR0pGpsVtTYi_yUdTGU-YUJ1UBqsHzLcAxaSDy5YDUEEOu8lCecSUk4RmEruEmFFDrg3ESNOlQDlBzKJFTW1HbtGjhhBlg52HjImAA7aFgUEEXJSWiUOnhgFBBUmHlofLnIABUgOzRYRJhYgHw6bLDJ1ABZbDmEZESkh8Q7PIC5PAIGXgZpNYA6vCg4eCx0jQU0urAa17h2VIi8nNejB4g77MC4mACnkDjAmEW9Kkgoivx4Nkkp_GhU2DSMFnzarIgV8ADsmwRFJxg7iQC7-ABKLSA6OSA1JLj4CGd5GJDAAMg6-GgEEHSY--gTpWwA1LgEBTXeVQg7EFw7cNy5GAA5ZP72GDvoXErgnGY8loAA7Jt82DvkLDsYwHSM-WAZKzgwOVQwSrgkN-xUzGWnlpCrZFBXiGSMhUoEFEUbp5x0gQYAutAcBLeFbHSMBEAETDYkVMzYjAOGEFSOhwTqRBx52DxFG4REO6A4dRkF6DqoXESPxFB0jYQt9DjZpAP24HuErDUtKpAI1LQ0jSoYEDq8JLgsXDjsUBDA7NtUsAcQBxw1GFRAZtD64Bw4jFgA2DiYWGSMiswwNRu1PLjsNtQINIy7_BBlJpW4OYCsRJiEhDm8IMlYDFksNYSA5shbpGg6HMh1PADEhcAkEGVIANw5UNQ6bSS4pABaNGzKVTAk4AUQuKQDl3A6tCXVlLt4rGcoWrjJBRRUpDd0uIAQIMTYyEQQRoU6TBQgzNTsNBBEmKRkWUCcZdaGADqktFSYO4CYJBB0mLiwBMWgFJRKvJi4pAOmABDsxJqEgFqckDvkaLg8EUj8BoQw-NAbV4g14Is0IGcROPBgSVwgWWwgdKSrPLw1PKXUhni69CC5ICREpEkceDkseMvAACDUxOw0EESkufTgZey2k9UdW3RgOFhsuKDlKxQNC0BGheKF7GWxhbA7tYRGSInIKGSP1MQ3eDuYLDmYoHSMSMRoEOzV1MCIkCxlGocQOahoRRg4KOg5bRx0jgQsOCw8VIw6dCwA7MuRJpfQqfCY18hlpTlcCZeQOmzAyHgNJkg5FKRFyJnAaLlsMLk0ELQQuZQAZe22DFcEaXycBCB0mLowfDU9OxwCFgoGGNTOJCw7UCS5dBCKLCQ1M5XAOKy8ZmAlMFtUfDSbhcw68Li4mAKHOCQQRJmkuFk0RHU9KuQMSrgkO2AsuFwIO2QguxggBdTrrBwQ2OQ7kIyqvC1KWEBIoHAA7JqYkgYcINDs5LiUEPhoCIvgPGbgiZyAtBM1ILgQBMvYOQeZF1hlDFlI3DUAO0xA-Ohjhoi4xCA4-CAEQGUMBEyEpEUYBDaEFLiMAARAuVAgJVjJDGBYxCBlGwZ3h-wVp_es6-zkVEAUoEqMPEnc6Duk6DddKSQTl2AA7ck9CDjEIDgA-EUY1bRlpIictDWkikC0ZI2FCNhcPDnsODhs5LuwBSooLqVkOFCwyJgESYDMAMA4WVBF1FvQiFpE1GXjpeTJQBy5vIx0pKp5fDccavCUBCC4pACL8NLXuLrQkHVIOFQkJBBVSDgIOFlMPLikACckyzQAF8gH2LncDme_9q8W3wcMuUgAJOw6kHRV7KmgVHaRGIhYWJSYOTxEucQEWsBYOtBYVUlY0HRb-Qg4APzkf6TIOuiYuewCe91kuihYtmhalHg4eHS4xBUkr5YwRKS1tAQguewCJ-sEAESkO0Aw6zSRKFRciuTg9QkpKKCp9Dx0pRqw4WpUfADUOKBMFBBGeVqkdEksRDu42LRYWJDcOlwsyaAE9-Q0pGogILjwBFtkXFtgeDSYuZSEuxwAOGwwFBBUpBdoOxxYyhAISbi8yJjYu6QIZUqGOCDI7NTm0FmwQDqc3HSYWzQkWIi8RoeX64YYyRAQqlj8RKekj4S8uUgClWw6hMZVEIsEPGZ4WKygSolwVT2nCPlUCEp8SAbcZKVYJAiIGMiYqESpEEh17fTgNyia7IDLKACoxGRUpEl0xYSouUgAlLg7KDBl7EksMDuc-MlIAEiEZDsgbIocMLjgLHaQS8DrhGjnAKhUNGSkODhAOZzcRoU6kBwA1EtkxVfYi_A0ZRkqUCxKJCA5SES4mAAWRBDsxGZVBshbsDR0oIlIJLTVZKC77BA77MQ5xSBEmDpgJBDU7MqQiIkUPDUlKtRASFAsOmEwNI8EZDjUUHY8SJAsqkC1KHggOhxEusAcOyw3B3R1GAVZhnxVpMk8QJkEIFggSADZV_BUQOUgiMQgNtEouUQETDooIESMWpT0yMRoSjwoAOyqCE0YIMhYqFA76DX5jC0rvMiL2NC61AbmXGVJWAWwWLy4OjA0RxBZ8Hg6tHDLQBCoXbC0QEpElxewuVgIukQAVKcX-QgIHJugcDSku-Bs5iC6yCBUpEnk1Dv8eMikADuEiBQR-KQASTBcONCXCKQAOcCYNBB17LuYlFaSFgBLEJBnNSgQdWlYdElsfOl8fCDA7MgkELlIADqESARYAMRnNAQ8JBDIpABK6HQFwEaQW2h8SJx8dKYnsEn1fFSkqHA0ZzYn_DgkMIhcJVlkZLiAZLR8aHAoBCC5SAEavBxb9Rw4BSDLsAUY-AgnNDqcgMvUFFps4Ep84DXsAMhLxXw71Xy57ADlxdc8aGEkBCB2kFvAeDv4bFfYW2SYOKycuSAG1wg14FgMgQloDFg4PDqATFSkqwQo5bhbrUA7vUBl4Eh9VDiNVLngALlAHDXtaFSehPhJHCNWZ6RESYBMde0okRyLNCzIrXS6-CA17Lgo_Gc0u1gERKRYvCRJZPx0pFr8MDssMFSkS33BCSAGdZg1SFm4qPpcBLkYIDSlt0QEILnsAhScO1AhVEggxNTgRBC4pAH0eEVLBeQkELikAaS0u_hkIMTI0EQQuKQBKxwl9Gz0fwe4JBBV7FgcOEpxDGSka9godymVTDmAMHSYOVQsS0x0RxyLvDS5KBCKFEA0mFrEZADguPAESFRMO8CUNIxKkHQA7MtVbEoUQKl4TgbHl3hm1PuoTDksTDoYNHSM-fQgh7zosIRLOEA4LKR2MEkIZGWlhCQ6cMBEjAWYS8RkdIw6IEC6pGA5SC06bEB5VWQ1uTloIDuwTKu8TDl88Dl5LHW4Oax0OLUoRRiWWOmwfDtAcBDs5FSMa7RYBCB1J5blBySoXPQ6IDwkELikASgsIIusLLnIBGv41AQgRnomxFnwOPTJOBiUepT0yUgAWwyMyxyMSFy4SdFIuUgAWWVIOZw8RexYND-FnLmkDFtENQU7d7E4wZbkseRrpETrvJgkWBUMRey4TABn2FpMnFpwITQNWmgZKiAQW-CxCzQBKv0dB5AkELnsAEuEnDuUnFc1awwYWh2oOxRARpFYiBhbiNg7TPh1SpVHB9C57ABabVwQ7MiZ1Ql7bLQ7UCAUEFXsW7C86yFfJIg7GEfUY6Xo-kAcOjwwywBkq_zU5l04fEWHzCQQuKQASzA826QGpyTr_GkpFAVriB0aaBhahC8HDLt8CSi5dLm0DHc0eMg8AOVnftaoyWgNK1QNSHwFJeg5UElnfaX8OChEuewAS6iDhFSZFDyrCEx2kEkI5DkY5maKRUDY0CIUUDnZ4ddUJE4UrLlIA_X1xWiq4FBl7IngYESYl0w6jETLpAUpkBw4pDwUELt8HTk0EoVIJBB14FtYxADQBCBV7ndsZoUo7AmnlPuAwJoAReS5pAWHaHVIm3w150lZGOSVqIW5ZOxKTEQ4aMC7SA0nMFrEqMUgiNAg2iwslqSG1FSkSogkO4A0ycQFFN0E7Fc0p6QllOR9OLwxS3wIOlUMSmEMNoSJrPR1M4aThpw0jIooQHSMOXgoOuSANIxYOGR0gDrEJDu4YDSAOqw8SyBgdIw4aCC4dCAEgDusYNiMALokADkAIBTNmIwAdiSq6EA4pPiosPhUQGUgSOggEOzd1eGIxHhZtCw3XSmMITpoQElAmFl4SGXIWRgwOixyZPxKtEA71HjJWAhL6GBIqTRF4GRMyKQBKxxwdEx17KpE-DcomKwsyCgmFujakJRLqDA4rCzLjBIm2LmYIFlktADMBCB176esyRwdaIwMWJxEOMxEVpFaQERIzCQ4SNxUpWscDic8OPS0RUi4-BjkfGggQAQgVKQWzDl89Lp4DiYESgXYZKQ7HCAAyAeguUgAOShSlJSJzDFbDEA0lAQgRUialFzI4BRYQCA6DDxUpEilIFnkWLs0ADncPBQQRKRn2LuwBFoowNrgIIoNcLmcCLuc4TT4WvS5Cwz9KVjgJig72Nx2kFhAYMi04FogODvg2LqQALokHMnsADkksLnEBJvVQtWEumAk5HyYCDFkVCXcBgx0pFu0JyW0RexbrLw73Ly4pABLFLw6XFTlxUkgBLi0IEVIFew4pCi57AGnuFvscESlFkGlHHc3hX8GpJlIYLqMCHSkqaUIVUkXeYTwdKQn2FkpBFSkSJAzh4zIpAA78XjYpABZmEQ7EcR0pUoYSEjsKDptxLikARoYSqSAS70EdKUm5ADcBCBGkLusDAOIijDIWVAwOpgwVKS48AC4pAIUmAUcRUlZIAS7FAzUfHtVnLq8DqSAyhg0JE6EzLj4CLjYFDVIuEwAdpBKDGQ6HGVVnCeBCzhMh7BbBChlSJmwzLlIAEqEODkALFSkWQBsO-VEuKQBKJybF4z4zLQ4vGQ71WRVPoQcWWxsdT4F0Dg1eESZhADqPECIcCDETDnwIDn8IHb4OKQguTwgREDIjAA6_MBFFFuIhMh8ABTURIBV1HSMOoQgOpAgNiBJ9EQ4OIQUjJnEIPhw3Dj83ADMO2zcFKBJsIT7LEUqUCCE_EsAjDW5OTR4quy8NJiLkEBlsFsgoNr40DhoRNh0RTroIEv8bDjEdMsMDEsUnDgeFNbRpmg66CDLDAxLUChJZhRGeyaU-0k55sFnNQbcWvgwuygABFjrNAgWKAY4ycQMqeAstGaUaEnkLLrkECRY63QsFswg7MTcypQYW9z7hkBVSFjoZDvVFLqQACaASBiIyewAJZT2XBWgW8wsNKembPuY3Su0HFjc_FjgNGVJSIBFSkAQSwAwphxF7UtgFSmcEFsMJDucILnIH6V8AN6FSFVISLXAOrEUuKQBS2w8OcxwFBB3NFtkJhaoRUi4zBwDitdgmiQ8mxRphmhb2DTIpABKwICEnEVJWpACp7uHtHc0IMDsyCQQuUgAFEg7iaRWkVu0HJt4_tQ4AMg5YDQkELlIAIgdAGSkWXSMOaSN2KQAWIQsWuRk9HyqXQi3sLlcOHSlGVApWQ0AiTBEmlzAu9ggZUk7tB0X1Pl0DSrkCLnwMGVIuaAANzSa-ai6vAy5LExEpJakhrTJIARL5CA79CDlxFssKOpJKFoAkYTQinAsS0TsOJBw27QcSBhIOkgxVZ2lwRuICJawuaRtaGiaB9wkEEc0WGBkOVQ4yPwgWqFgSoxgN9gA5DikzADs2xxEWLSQO9Q0RJgAxEg9cDpwwLkUBOa01HBbqHBYeYTmXJqUMvR0OnQhGFxsOQCMFBBV7BcYByi6XARbbGjJRIiL-Lxl4IqE5De1K1wYOUwsSTwgNIyLzMBlGElUQKjwIDskIoTYuIwAWZinZnw5_CFKCCJUJDmAIMrwGADIS8ilVFKEtADcOGDsZiT6aGQ4DMRIpRwUjJmAIDnktEr1HDdQOUy0SOxoFKBZgCBbEbiL_DCJwCBkjDjEZDl9rEUZKZk8-kB0SaC4AOzIwFSKtGQ2MwZUSvEEZaRIFCgQ7OFVETsJdEmYIMg4RShUqLn0CDW8WmycOqW0yPAZOXxxFdhJfVxmbEjMNDjcNJjMNbl8cDisadTQlvjbBASZwEVXfEu4iADsyBU8W0Q0O0CrVX1XJLpsAZT0OYB4V7QU1CTkdKXX1VZ8auxwyEwFptSFEFU8SkAsOKQs26wIBYQUEFSlBvAkEHXguCQgxjoX3gfsuzAMuJ3gNKV61BQ6ZHwUEFSkWDgk--EAODiUJBBEpIrhmPRkW3CiBnhEmVi4KxS3hyyZUOqnuPk4RJqkKJtIKFqMIATEdexbpEA4QHTVuJtwJLmMFLsoZMRwWBnIOEnIyUgAAOR4zCBUpwXsJBB0pLgcODVLJkkKYGCreDREpFqUKDmUWLlIAFuoWDh1LFSkWaA4BbR0pFmAbADcyBQoSIhkOJhkuJAsu1gENeyaACjJ7AEaLD1YgGDXZJllGFuYIFrQkAOIiLgoWlAgBzVVkFjEhPr4MLicdDaQVZTJsDBJPJMExJk8kXrQPBY2BdRlSDkRoPtMQGp8gEuojEXvpneGpLs0AKuAVESkJTg6TDS4pABa9Hg6EIhEpGhkILrENBDc5DhAYJuAIEmMfADsyy1IO2igSuWENb0rfBBILOgQ7NjVfIg4PPYUORysJBBFJxYsAOzLlHg45KQ65YRUjDqIXADsyZSflHw4lOA2PSmwADlwpLi0HIkVHGY8VEw1GEggQDgIZGSMSRAgqEiEOAQgSCxAZI6Eh4eQVRhY0CC4zBYGzEsUpDWkOuzUAMA7lKRlGPmU15dkOFyIdIwAwCTkRRhI_EdWJJicIih0BPlAzDqcQOs9lQm4AgZoOnUkZkeG9Dt0KVpEAEpAIFccR-kZqjiJAQREjHhhxGWk-ByIOTxYOThYdIw5-KzIoZSE9NkABDvweLlgCSiMADrIREhRJDYwiRgoZjD6tCU3QAQgdJmUtBDs1NRtWKhMO2xEOT2YRbxoWZwEIHUxO7jISrzgOszgyawoWiygOSSYVUhK8CEJABEooCZ0qGe0u0B8tE06XKxamFg4-NxEmSlUKSrUEGiA2MsQA1Z45FhZ-Gw4cDC6bABZSFA6NLRF1Lq4AHcRKvh4SLiEOMiEydQYiIA0Vxy5QEh1SSjMVEpc3Dps3Ms0AxXHBdZXNLl4hGVIu0CstaCbfCTL6BSLhEiaKCFbvJsViwWYZe1ZBJxKHLw75WrWUwQAENTsyAAZJSA4ICzUcYU8STiIZxy4pBw3HSvImLkATDSZOxgkiyB6V8SJANh1ySnsCyS8uHQRlfqGC_asSdgwO5BQyVQJpVBL6Fg2bLpECHXgqCQkRKRbfEA5qCC4pABJhCCGINTkuJgEdUgWRMgsNtRodJiIEK9VwIkcnGSYhvgQ0O_2LDq9dNjFXBDYwCTkRvuEMBDU7MvIbPuMNSpooIsVPLQQOrQoSvjAZjBU2DSNK3wU-onihdhKHKBlGDtkTBDA7JoE4IbphWB0j4UIOJFcRaRKHCw6LQhlGQQMOZh0VI8FFDkoxHSMOOAwqskBORgAOAhAuh1IBUw4TWBlGAVkODBwRaeUnDlFSHSMO2hMqjUgh-w6cER0jFVktOw6NCDkY-fEF9yG8DShKYwFNArUrLS0yQAE6XwgVVhm0QSISEFAdaeHRLiMA8b4RIw44EDY7ECUaIXYNI0qIBz6GASKoJBmMFTYNRhbOCyGbLjAFBDgwgTYmJCAWTyZGRyMSnEEOaUB2TwAWYR4OAQp1f1rIEhIMPQ7nO1kEQbEJBB3HSssZLtQEGfASEybhpSZ0Flp1Bgl3DhwkEc1KywGBIy6uCxIsCwA7MhAHPt0DJr4mMjYH_Z4xixZdEA4KQh3BItkRESYWvhmhDx0mLvcGESkWAxUOByAuKQBhhw6MCyqTLxIXDQ4EDS6SAiKZGSbENC7qBz1iSkwGDq4-BDs3LpUDxfQOez4RnhZiGwQ7MTIYGRYHYTKvUBU5HXIOLAwJBBFPTjkJgQgOnRTZvkomAEnqJsUNSjMD9RoRuxJ_FQ68DC4tBw6VCg5peBEmSkkAPhoGBTAAOzJscEJsAA5dDQkEPQGBhy47eiIMaDkkNYMNtRJ_OwQ7ODJcGmn3gQMRJib1CS7jBEobA1LPChKWFmHaEU8OkR0-dQAluA6MNjW8DmInOltTlf8NviHLOiwGgewFEw0jFTM5BCUkBDszNZMOCDs6gkMahDciLAkSxwxG60MOuXERaYGMDuUlLmkADiYIDikIEYwOFQ4OQWYdjA4gDQETDSNKIQcEMjkWPg4RIwFmOv4GBRDhEQ1GFXkZaeUO_WfVNx0jDis34XpmRgDtZxWMDRAywgTh8P3zIRgO3h0daSL9Dg1pSkYAFZ9mIwBB_6EWESNBPjAzOzE0beKWhBtbMG0K"
)

var staticFiles = map[string]string{
	"/cat.txt": c_CatTxt,
}

// Lookup returns the bytes associated with the given path, or nil if the path was not found.
func Lookup(path string) []byte {
	s, ok := staticFiles[path]
	if ok {
		d, err := base64.URLEncoding.DecodeString(s)
		if err != nil {
			log.Print("main.Lookup: ", err)
			return nil
		}
		r, err := snappy.Decode(nil, d)
		if err != nil {
			log.Print("main.Lookup: ", err)
			return nil
		}
		return r
	}
	return nil
}

// ServeHTTP serves the stored file data over HTTP.
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	if strings.HasSuffix(p, "/") {
		p += "index.html"
	}
	b := Lookup(p)
	if b != nil {
		mt := mime.TypeByExtension(path.Ext(p))
		if mt != "" {
			w.Header().Set("Content-Type", mt)
		}
		w.Header().Set("Expires", time.Now().AddDate(0, 0, 1).Format(time.RFC1123))
		w.Write(b)
	} else {
		http.NotFound(w, r)
	}
}

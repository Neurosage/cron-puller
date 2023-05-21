import pandas as pd
import os
import gspread
import subprocess
import argparse

def pull_and_update(days=0, bio=False):
    subprocess.run("cd api && go run puller.go -d {}".format(days), shell=True)

    gc = gspread.oauth()

    sh = gc.open("Weight")

    data_loc = './data/'

    biometrics = "biometrics.csv"
    nutrition = "nutrition.csv"

    nut_df = pd.read_csv(data_loc + nutrition)

    cals = "Energy (kcal)"
    date = "Date"

    nut_df = nut_df[[date, cals]]
    # for all dates , change format to M/DD/YY
    for i in range(len(nut_df[date])):
        d = nut_df[date][i]
        [y,m,d] = d.split('-')

        #strip leading 0 from month
        if m[0] == '0':
            m = m[1]
        
        #turn year from 20XX to XX
        y = y[2:]
        nut_df.loc[i, date] = m + '/' + d + '/' + y

    # for all dates, find the cell in the spreadsheet
    for i in range(len(nut_df[date])):
        d = nut_df[date][i]
        c = sh.sheet1.find(d)
        
        #put the calories in the cell
        # if c is not None:
        print(c)
        if c is not None:
            sh.sheet1.update_cell(c.row, c.col + 3, nut_df[cals][i])

    # biometrics
    if(bio):
        
        bio_df = pd.read_csv(data_loc + biometrics)

        metric = "Metric"
        bio_date = "Day"
        bio_val = "Amount"

        bio_df = bio_df[[bio_date, metric, bio_val]]

        # only when metric = weight
        bio_df = bio_df[bio_df[metric] == "Weight"]

        # for all dates , change format to M/DD/YY
        for i in range(len(bio_df[bio_date])):
            d = bio_df[bio_date][i]
            [y,m,d] = d.split('-')

            #strip leading 0 from month
            if m[0] == '0':
                m = m[1]
            
            #turn year from 20XX to XX
            y = y[2:]
            bio_df.loc[i, bio_date] = m + '/' + d + '/' + y

        # for all dates, find the cell in the spreadsheet
        for i in range(len(bio_df[bio_date])):
            d = bio_df[bio_date][i]
            c = sh.sheet1.find(d)

            if (c is not None):
                sh.sheet1.update_cell(c.row,c.col+1, bio_df[bio_val][i])




def init_argparse() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "--days", "-d", help="Number of days to pull", type=int, default=0)
    parser.add_argument(
        "--bio", "-b", help="Put biometrics in the sheet", type=bool, default=False)
    
    return parser

    


def main():
    parser = init_argparse()
    args = parser.parse_args()
    pull_and_update(args.days, args.bio)

if __name__ == "__main__":
    

    main()
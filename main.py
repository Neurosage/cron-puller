import pandas as pd
import os
import gspread
import subprocess
import argparse

def pull_and_update(days=0):
    subprocess.run("cd api && go run puller.go -d {}".format(days), shell=True)

    gc = gspread.service_account()

    sh = gc.open("Weight")

    data_loc = './data/'
    ls = os.listdir(data_loc)

    df = pd.DataFrame()
    df = pd.read_csv(data_loc + ls[0])

    cals = "Energy (kcal)"
    date = "Date"

    df = df[[date, cals]]
    # for all dates , change format to M/DD/YY
    for i in range(len(df[date])):
        d = df[date][i]
        [y,m,d] = d.split('-')

        #strip leading 0 from month
        if m[0] == '0':
            m = m[1]
        
        #turn year from 20XX to XX
        y = y[2:]
        df.loc[i, date] = m + '/' + d + '/' + y

    # for all dates, find the cell in the spreadsheet
    for i in range(len(df[date])):
        d = df[date][i]
        c = sh.sheet1.find(d)
        
        #put the calories in the cell
        # if c is not None:
        print(c)
        if c is not None:
            sh.sheet1.update_cell(c.row, c.col + 3, df[cals][i])

def init_argparse() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "--days", "-d", help="Number of days to pull", type=int, default=0)
    
    return parser

    


def main():
    parser = init_argparse()
    args = parser.parse_args()
    if args.days:
        pull_and_update(args.days)
    else:  
        pull_and_update()

if __name__ == "__main__":
    

    main()
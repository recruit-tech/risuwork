"use client";

import {createTheme, ThemeProvider} from "@mui/material/styles";
import CssBaseline from "@mui/material/CssBaseline";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Container from "@mui/material/Container";
import Grid from "@mui/material/Grid";
import Paper from "@mui/material/Paper";
import Footer from "@/app/components/Footer";
import Header from "@/app/components/Header";
import Chart from "@/app/components/Chart";
import Ranking from "@/app/components/Ranking";
import {useEffect, useState} from "react";
import {ChartResults, fetch, RankingResults} from "@/services/data/benchResults";

const defaultTheme = createTheme();

export default function Dashboard() {
  const [chartResults, setChartResults] = useState<ChartResults>([]);
  const [rankingResults, setRankingResults] = useState<RankingResults>([]);
  useEffect(() => {
    (async () => {
      const dataSet = await fetch();
      setChartResults(dataSet.chartResults);
      setRankingResults(dataSet.rankingResults);
    })();
  }, []);
  return (
    <ThemeProvider theme={defaultTheme}>
      <Box sx={{display: "flex"}}>
        <CssBaseline/>
        <Header/>
        <Box
          component="main"
          sx={{
            backgroundColor: (theme) =>
              theme.palette.mode === "light"
                ? theme.palette.grey[100]
                : theme.palette.grey[900],
            flexGrow: 1,
            height: "100vh",
            overflow: "auto",
          }}
        >
          <Toolbar/>
          <Container maxWidth="lg" sx={{mt: 4, mb: 4}}>
            <Grid container spacing={3}>
              <Grid item xs={12}>
                <Paper sx={{p: 2, display: "flex", flexDirection: "column", height: 400}}>
                  <Chart data={chartResults}/>
                </Paper>
              </Grid>
              <Grid item xs={12}>
                <Paper sx={{p: 2, display: "flex", flexDirection: "column"}}>
                  <Ranking data={rankingResults}/>
                </Paper>
              </Grid>
            </Grid>
            <Footer sx={{pt: 2}}/>
          </Container>
        </Box>
      </Box>
    </ThemeProvider>
  );
}
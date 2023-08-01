-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Aug 01, 2023 at 04:43 PM
-- Server version: 10.4.21-MariaDB
-- PHP Version: 8.1.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `dbo`
--

-- --------------------------------------------------------

--
-- Table structure for table `dbo_customer`
--

CREATE TABLE `dbo_customer` (
  `id` int(11) NOT NULL,
  `nama` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `email` varchar(50) NOT NULL,
  `nomor_hp` varchar(100) NOT NULL,
  `jenis_kelamin` varchar(10) NOT NULL,
  `tanggal_lahir` varchar(100) NOT NULL,
  `kota_domisili` varchar(100) NOT NULL,
  `tgl_input` datetime NOT NULL DEFAULT current_timestamp(),
  `status_akun` int(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `dbo_customer`
--

INSERT INTO `dbo_customer` (`id`, `nama`, `password`, `email`, `nomor_hp`, `jenis_kelamin`, `tanggal_lahir`, `kota_domisili`, `tgl_input`, `status_akun`) VALUES
(1, 'Jemi Yosua Laoere', '123456', 'jemiyosua@gmail.com', '082118009042', 'L', '24/01/1998', 'Jakarta Barat', '2023-07-31 22:31:03', 1),
(2, 'Yosua', '123456', 'yosua@gmail.com', '0821180090421', 'L', '24/02/1998', 'Jakarta Barat', '2023-07-31 22:34:00', 1);

-- --------------------------------------------------------

--
-- Table structure for table `dbo_login`
--

CREATE TABLE `dbo_login` (
  `id` int(11) NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `tgl_input` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `dbo_login`
--

INSERT INTO `dbo_login` (`id`, `username`, `password`, `tgl_input`) VALUES
(1, 'jemiyosua', '123456', '2023-07-30 22:52:05');

-- --------------------------------------------------------

--
-- Table structure for table `dbo_log_error`
--

CREATE TABLE `dbo_log_error` (
  `id` int(11) NOT NULL,
  `page` varchar(100) NOT NULL,
  `body_json` longtext NOT NULL,
  `query` longtext NOT NULL,
  `json_return` longtext NOT NULL,
  `errorcode` varchar(5) NOT NULL,
  `errorcode_return` varchar(5) NOT NULL,
  `errormessage` varchar(1000) NOT NULL,
  `errormessage_return` varchar(1000) NOT NULL,
  `source` varchar(100) NOT NULL,
  `tgl_input` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `dbo_customer`
--
ALTER TABLE `dbo_customer`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `dbo_login`
--
ALTER TABLE `dbo_login`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `dbo_log_error`
--
ALTER TABLE `dbo_log_error`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `dbo_customer`
--
ALTER TABLE `dbo_customer`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `dbo_login`
--
ALTER TABLE `dbo_login`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `dbo_log_error`
--
ALTER TABLE `dbo_log_error`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

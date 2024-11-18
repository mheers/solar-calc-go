# monthly flux has band 1 to 12 and comes in 32bit float - ok, but colours are missing
gdal_translate -b 6 -of PNG -scale ../results/VHJpZnR3ZWcgMTUsIFBhZGVyYm9ybg==_monthlyFlux.tiff _monthlyFlux.png

# dsm has band 1 and comes in 32bit float - ok, but colours are missing
gdal_translate -b 1 -of PNG -scale ../results/VHJpZnR3ZWcgMTUsIFBhZGVyYm9ybg==_dsm.tiff dsm.png

# annual flux has band 1 and comes in 32bit float - not working yet - only black
gdalwarp -t_srs EPSG:3857 ../results/VHJpZnR3ZWcgMTUsIFBhZGVyYm9ybg==_annualFlux.tiff annualFlux_reprojected.tif
gdal_translate -b 1 -of PNG -scale ./annualFlux_reprojected.tif annualFlux.png

# mask - not working yet - only black
gdal_translate -b 1 -of PNG -scale ../results/VHJpZnR3ZWcgMTUsIFBhZGVyYm9ybg==_mask.tiff mask.png

# rgb - ok
gdal_translate -b 1 -of PNG -scale ../results/VHJpZnR3ZWcgMTUsIFBhZGVyYm9ybg==_rgb.tiff rgb.png
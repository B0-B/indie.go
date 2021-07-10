import matplotlib.pyplot as plt

sizes = [50, 250, 500, 750, 1052, 2000, 5000]
times = [.208, .271, .413, .602, .886, 3.54, 9.52]
caps = [1.25, 31.3, 125, 281, 553, 2000, 12500]
originals = [4.6, 87.8, 412, 795, 1600, 4700, 22700]
relativeStorage = [100*caps[i]/originals[i] for i in range(len(caps))]

# plot
fig, axs = plt.subplots(3)
axs[0].plot(sizes, times, label=f'Times [s]')
axs[1].plot(sizes, caps, label=f'Capacitiy [KB]')
axs[2].plot(sizes, relativeStorage, label='Relative Storage [%]')
axs[0].legend()
axs[1].legend()
axs[2].legend()
plt.xlabel('Original Image Height [px]')
plt.show()

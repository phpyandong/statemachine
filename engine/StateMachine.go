package engine
/**
  private final GlobalContext globalContext;

    private final StateConfiguration[] machineConfiguration;

    private final ScheduledExecutorService timer = new ScheduledThreadPoolExecutor(1);
    private Future<?> delayTransitionTask;

    private State currentState;
    private StateConfiguration currentStateConfiguration;

    /**
     * 触发状态变更
     * @param trigger
     * @return
     * @throws UnknownTriggerException
     * @throws MissingStateConfigurationException
*/
public synchronized State fire(final Event trigger) throws UnknownTriggerException, MissingStateConfigurationException {
/**
 * 获取下一状态
 */
currentState = currentStateConfiguration.fire(trigger, machineConfiguration,globalContext);
/**
 * 获取下一状态的配置信息
 */
currentStateConfiguration = machineConfiguration[currentState.getPosition()];

cleanFormerDelayTransition();
/**
 * 开始延时传输
 */
startDelayTransition(currentStateConfiguration);
return currentState ;
}
 */